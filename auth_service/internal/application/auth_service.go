package application

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/auth_service/internal/domain/service"
	"hype-casino-platform/pkg/db"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
	"hype-casino-platform/pkg/pb/gen/auth"
	"hype-casino-platform/pkg/req_analyzer"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	authService   *service.AuthService
	clientService *service.ClientService
	db            db.Database
	reqAnalyzer   req_analyzer.ReqAnalyzer
}

func NewAuthService(
	authService *service.AuthService,
	clientService *service.ClientService,
	db db.Database,
	reqAnalyzer req_analyzer.ReqAnalyzer) *AuthService {
	return &AuthService{
		authService:   authService,
		clientService: clientService,
		db:            db,
		reqAnalyzer:   reqAnalyzer,
	}
}

func (s *AuthService) ClientAuth(ctx context.Context, req *auth.ClientAuthRequest) (res *auth.AuthResponse, err error) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Begin transaction
	ctx, kgsErr := s.db.Begin(ctx)
	if kgsErr != nil {
		return nil, kgsErr
	}

	defer func() {
		// If there is an error, rollback the transaction
		if err != nil {
			_, rollbackErr := s.db.Rollback(ctx)
			if rollbackErr != nil {
				kgsotel.Error(ctx, rollbackErr.Error())
				err = rollbackErr
			}
			return
		}

		// Commit the transaction
		_, commitErr := s.db.Commit(ctx)
		if commitErr != nil {
			kgsotel.Error(ctx, commitErr.Error())
			err = commitErr
		}
	}()

	// Create client token
	result, kgsErr := s.authService.CreateClientToken(ctx, req.ClientId)
	if kgsErr != nil {
		return nil, kgsErr
	}

	return &auth.AuthResponse{
		AccessToken:     result.Token,
		TokenExpireSecs: int64(result.TokenExpireSecs),
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Begin transaction
	ctx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	// Get user ip and user agent info
	ipInfo := s.reqAnalyzer.GetIpInfo(ctx, req.Ip)
	userAgentInfo := s.reqAnalyzer.GetUserAgentInfo(ctx, req.UserAgent)
	defer func() {
		// TODO: Check user location and device, if unsafe, send notification to user
		// Get last login record
	}()

	// Login
	result, loginErr := s.authService.Login(ctx, req.AccessToken, req.UserId, req.Password)
	// If Login failed and the error is not a password error or account locked error, rollback the transaction
	if loginErr != nil &&
		loginErr.Code().Int() != kgserr.AccountPasswordError &&
		loginErr.Code().Int() != kgserr.AccountLocked {
		_, rollbackErr := s.db.Rollback(ctx)
		if rollbackErr != nil {
			kgsotel.Error(ctx, rollbackErr.Error())
			loginErr = rollbackErr
		}
		return nil, loginErr
	}

	// Stone login record
	errMsg := ""
	if loginErr != nil {
		errMsg = loginErr.Error()
	}
	loginRecord := &entity.LoginRecord{
		Browser:     userAgentInfo.Browser,
		BrowserVer:  userAgentInfo.BrowserVer,
		Ip:          ipInfo.Ip,
		Os:          userAgentInfo.OS,
		Platform:    userAgentInfo.Platform,
		Country:     ipInfo.Country,
		CountryCode: ipInfo.CountryCode,
		City:        ipInfo.City,
		Asp:         ipInfo.Asp,
		IsMobile:    userAgentInfo.IsMobile,
		IsSuccess:   loginErr == nil,
		ErrMessage:  errMsg,
	}
	_, err = s.authService.AddLoginRecord(ctx, req.UserId, loginRecord)
	if err != nil {
		_, rollbackErr := s.db.Rollback(ctx)
		if rollbackErr != nil {
			kgsotel.Error(ctx, rollbackErr.Error())
			err = rollbackErr
		}
		return nil, err
	}

	// Commit the transaction
	_, commitErr := s.db.Commit(ctx)
	if commitErr != nil {
		kgsotel.Error(ctx, commitErr.Error())
		return nil, commitErr
	}

	// If login failed, return the error
	if loginErr != nil {
		return nil, loginErr
	}

	return &auth.AuthResponse{
		AccessToken:     result.Token,
		TokenExpireSecs: int64(result.TokenExpireSecs),
	}, nil
}

func (s *AuthService) ValidToken(ctx context.Context, req *auth.ValidTokenRequest) (res *auth.ValidTokenResponse, err error) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	res = &auth.ValidTokenResponse{}

	// Validate token
	payload, kgsErr := s.authService.ValidateToken(ctx, req.AccessToken)
	if kgsErr != nil {
		return nil, kgsErr
	}

	res.ClientId = payload.ClientId
	res.MerchantId = payload.MerchantId
	res.UserAccount = payload.Account
	res.UserId = payload.UserId

	if payload.RoleId == nil {
		return res, nil
	}

	role, kgsErr := s.clientService.FindRole(ctx, payload.ClientId, *payload.RoleId)
	if kgsErr != nil {
		if kgserr.ResourceNotFound == kgsErr.Code().Int() {
			return res, nil
		}
		return nil, kgsErr
	}

	res.Role = &auth.Role{
		RoleId:   role.Id,
		RoleName: role.Name,
		PermIds:  role.GetPermissionIds(),
	}

	return res, nil
}
