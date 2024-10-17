package service

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/aggregate"
	"hype-casino-platform/auth_service/internal/domain/repository"
	"hype-casino-platform/auth_service/internal/domain/vo"
	"hype-casino-platform/pkg/kgscrypto"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
)

type UserService struct {
	clientRepo repository.ClientRepo
	userRepo   repository.UserRepo
	crypto     kgscrypto.KgsCrypto
}

func NewUserService(clientRepo repository.ClientRepo, userRepo repository.UserRepo) *UserService {
	return &UserService{
		userRepo:   userRepo,
		clientRepo: clientRepo,
		crypto:     kgscrypto.New(),
	}
}

func (u *UserService) CreateUser(ctx context.Context, clientId int64, userInfo vo.UserInfo) (*aggregate.User, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Find client
	_, err := u.clientRepo.Find(ctx, clientId)
	if err != nil {
		return nil, err
	}

	// Check the parameters
	if userInfo.Account == "" {
		err = kgserr.New(kgserr.InvalidArgument, "account is required")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	if userInfo.Id == 0 {
		err = kgserr.New(kgserr.InvalidArgument, "id is required")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	if userInfo.Status == 0 {
		err = kgserr.New(kgserr.InvalidArgument, "status is required")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Password could be empty ,if not empty, hash with Md5 , and encode with base64
	if userInfo.Password != "" {
		hashPwd, err := u.crypto.HashPassword(ctx, userInfo.Password)
		if err != nil {
			return nil, err
		}
		userInfo.Password = hashPwd
	}

	// Create user
	user := &aggregate.User{
		Id:       userInfo.Id,
		Account:  userInfo.Account,
		Status:   userInfo.Status,
		Password: userInfo.Password,
	}

	user, err = u.userRepo.Create(ctx, clientId, user)

	return user, err
}

func (u *UserService) UpdateUser(ctx context.Context, userInfo vo.UserInfo) (*aggregate.User, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Find user
	user, err := u.userRepo.Find(ctx, userInfo.Id)
	if err != nil {
		return nil, err
	}

	// If account is not empty, update it
	if userInfo.Account == "" || userInfo.Status == 0 {
		err = kgserr.New(kgserr.InvalidArgument, "user account and status are required")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	user.Account = userInfo.Account
	user.Status = userInfo.Status

	if userInfo.Password != "" {
		secretByte := u.crypto.HashMD5(ctx, userInfo.Password)
		user.Password = u.crypto.EncodeBase64(ctx, secretByte)
	}

	// Update user
	user, err = u.userRepo.Update(ctx, user)

	return user, err
}

func (u *UserService) GetUser(ctx context.Context, userId int64) (*aggregate.User, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Find user
	user, err := u.userRepo.Find(ctx, userId)

	return user, err
}
