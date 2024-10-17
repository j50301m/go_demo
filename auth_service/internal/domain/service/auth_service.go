package service

import (
	"context"
	"fmt"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/auth_service/internal/domain/repository"
	"hype-casino-platform/auth_service/internal/domain/vo"
	"hype-casino-platform/auth_service/internal/infrastructure/token_helper"
	"hype-casino-platform/pkg/db"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgscrypto"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
	"time"
)

// AuthService handles authentication-related operations.
type AuthService struct {
	clientRepo  repository.ClientRepo
	userRepo    repository.UserRepo
	tokenHelper token_helper.TokenHelper
	cache       db.Cache
	crypto      kgscrypto.KgsCrypto
}

const TokenPrefix = "token"

func NewAuthService(
	clientRepo repository.ClientRepo,
	userRepo repository.UserRepo,
	cache db.Cache,
	helper token_helper.TokenHelper) *AuthService {
	return &AuthService{
		clientRepo:  clientRepo,
		userRepo:    userRepo,
		tokenHelper: helper,
		cache:       cache,
		crypto:      kgscrypto.New(),
	}
}

// CreateClientToken creates a client token for the given client ID.
func (a *AuthService) CreateClientToken(ctx context.Context, clientId int64) (*vo.Token, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get client info
	client, err := a.clientRepo.Find(ctx, clientId)
	if err != nil {
		return nil, err
	}

	// Check client ia active
	if !client.Active {
		err = kgserr.New(kgserr.ClientInactive, fmt.Sprintf("Client id: %v is not active", client.Id))
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Create token
	payload := vo.NewTokenPayload(client.MerchantId, client.Id)
	token, err := a.tokenHelper.Create(ctx, client.Secret, payload.ToMap())
	if err != nil {
		return nil, err
	}

	// Cache token
	key := fmt.Sprintf("%s:%s", TokenPrefix, token)
	err = a.cache.Set(
		ctx,
		key,
		token,
		time.Second*time.Duration(client.TokenExpireSecs),
	)
	if err != nil {
		return nil, err
	}

	return &vo.Token{
		Token:           token,
		TokenExpireSecs: client.TokenExpireSecs,
	}, nil
}

// Login authenticates a user with the provided token (containing 'cid'), user ID, and password.
// It returns a new token upon successful login.
// When login is successful , the old token is going to delete from cache.
// Regardless of success or failure, a login record is created.
func (a *AuthService) Login(ctx context.Context, token string, userId int64, password string) (*vo.Token, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Find user by id
	user, err := a.userRepo.Find(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Validate client token
	calims, err := a.GetTokenPayload(ctx, token)
	if err != nil {
		return nil, err
	}

	// Get client, this is going to find from cache first and then from database
	client, err := a.clientRepo.Find(ctx, calims.ClientId)
	if err != nil {
		return nil, err
	}

	// Validate token
	_, err = a.tokenHelper.Validate(ctx, token, client.Secret)
	if err != nil {
		return nil, err
	}

	// Check client is active
	if !client.Active {
		err = kgserr.New(kgserr.ClientInactive, fmt.Sprintf("Client id: %v is not active", client.Id))
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Check user status
	if user.Status != enum.UserStatusType.Active {
		err = kgserr.New(kgserr.AccountLocked, fmt.Sprintf("User id: %v is not active", user.Status))
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Check user password is correct
	if !a.crypto.CompareHashAndPassword(ctx, user.Password, password) {
		// Increase login error count
		user.PasswordFailTimes += 1

		// If password failed time is more than client login failed times, then lock user
		if user.PasswordFailTimes >= client.LoginFailedTimes {
			user.Status = enum.UserStatusType.Locked
		}

		// Create password error with remaining times
		err := kgserr.New(kgserr.AccountPasswordError, "Invalid password").
			WithData(map[string]interface{}{
				"remaining_times": client.LoginFailedTimes - user.PasswordFailTimes,
			})
		kgsotel.Error(ctx, err.Error())

		// Update user
		_, updateErr := a.userRepo.Update(ctx, user)
		if updateErr != nil {
			kgsotel.Error(ctx, updateErr.Error())
			return nil, err
		}

		return nil, err
	}

	// PasswordFailedTime is reset to 0
	user.PasswordFailTimes = 0

	// Get user role.
	// If user role not found(no role) it just continue to create token
	role, loginErr := user.Role(ctx)
	if loginErr != nil && loginErr.Code().Int() != kgserr.ResourceNotFound {
		kgsotel.Warn(ctx, loginErr.Error())
		return nil, nil
	}

	// Create new token
	opts := []vo.TokenPayloadOption{
		vo.WithUserId(user.Id),
		vo.WithAccount(user.Account),
	}
	if role != nil {
		opts = append(opts, vo.WithRoleId(role.Id))
	}
	payload := vo.NewTokenPayload(client.MerchantId, client.Id, opts...)
	newToken, loginErr := a.tokenHelper.Create(ctx, client.Secret, payload.ToMap())
	if loginErr != nil {
		return nil, loginErr
	}

	// Cache token
	key := fmt.Sprintf("%s:%s", TokenPrefix, newToken)
	loginErr = a.cache.Set(ctx, key, newToken, time.Second*time.Duration(client.TokenExpireSecs))
	if loginErr != nil {
		return nil, loginErr
	}

	// Delete old token from cache
	oldKey := fmt.Sprintf("%s:%s", TokenPrefix, token)
	loginErr = a.cache.Delete(ctx, oldKey)
	if loginErr != nil {
		return nil, loginErr
	}

	return &vo.Token{
		Token:           newToken,
		TokenExpireSecs: client.TokenExpireSecs,
	}, nil
}

// ValidateToken validates the given token.
func (a *AuthService) ValidateToken(ctx context.Context, token string) (*vo.TokenPayload, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get token payload
	calims, err := a.GetTokenPayload(ctx, token)
	if err != nil {
		return nil, err
	}

	// Get client, this is going to find from cache first and then from database
	client, err := a.clientRepo.Find(ctx, calims.ClientId)
	if err != nil {
		return nil, err
	}

	// Validate token
	_, err = a.tokenHelper.Validate(ctx, token, client.Secret)
	if err != nil {
		return nil, err
	}

	// Check client is active
	if !client.Active {
		err = kgserr.New(kgserr.ClientInactive, fmt.Sprintf("Client id: %v is not active", client.Id))
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Check token is expired , if it expired can not found in cache
	key := fmt.Sprintf("%s:%s", TokenPrefix, token)
	_, err = a.cache.Get(ctx, key)
	if err != nil {
		err = kgserr.New(kgserr.TokenExpired, "Token is expired")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	return &calims, nil
}

// GetTokenPayload extracts the payload information from the given token.
//
// Parameters:
//   - ctx: context for tracing and potential cancellation
//   - token: the token string to be parsed
//
// Returns:
//   - vo.TokenPayload: the extracted token payload
//   - *kgserr.KgsError: an error if parsing fails; nil on success
//
// Security note: Ensure that the token's legitimacy and validity are verified
// elsewhere before relying on this payload for sensitive operations.
func (a *AuthService) GetTokenPayload(ctx context.Context, token string) (vo.TokenPayload, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	tp := vo.TokenPayload{}

	// Get payload from token
	claims, err := a.tokenHelper.GetPayload(ctx, token)
	if err != nil {
		return tp, err
	}

	// Get TokenPayload from claims
	tp, err = vo.ToTokenPayload(ctx, claims)
	if err != nil {
		return tp, err
	}

	return tp, nil
}

func (a *AuthService) AddLoginRecord(ctx context.Context, userId int64, record *entity.LoginRecord) (*entity.LoginRecord, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	return a.userRepo.AddLoginRecord(ctx, userId, record)
}
