package application

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/service"
	"hype-casino-platform/auth_service/internal/domain/vo"
	"hype-casino-platform/pkg/db"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgsotel"
	"hype-casino-platform/pkg/pb/gen/auth"
)

type UserService struct {
	auth.UnimplementedUserServiceServer
	userService *service.UserService
	db          db.Database
}

func NewUserService(userService *service.UserService, db db.Database) *UserService {
	return &UserService{
		userService: userService,
		db:          db,
	}
}

func (u *UserService) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (res *auth.Empty, err error) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Begin transaction
	ctx, kgsErr := u.db.Begin(ctx)
	if kgsErr != nil {
		return nil, kgsErr
	}

	defer func() {
		// If there is an error, rollback the transaction
		if err != nil {
			_, rollbackErr := u.db.Rollback(ctx)
			if rollbackErr != nil {
				kgsotel.Error(ctx, rollbackErr.Error())
				err = rollbackErr
			}
			return
		}

		// Commit the transaction
		_, commitErr := u.db.Commit(ctx)
		if commitErr != nil {
			kgsotel.Error(ctx, commitErr.Error())
			err = commitErr
		}
	}()

	// Convert user status
	userStatus, kgsErr := enum.UserStatusFromInt(int(req.Status))
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	userInfo := vo.UserInfo{
		Id:       req.Id,
		Account:  req.Account,
		Password: req.Password,
		Status:   userStatus,
	}

	// Create user
	_, kgsErr = u.userService.CreateUser(ctx, req.ClientId, userInfo)
	if kgsErr != nil {
		return nil, kgsErr
	}

	return &auth.Empty{}, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *auth.UpdateUserRequest) (res *auth.Empty, err error) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Begin transaction
	ctx, kgsErr := u.db.Begin(ctx)
	if kgsErr != nil {
		return nil, kgsErr
	}

	defer func() {
		// If there is an error, rollback the transaction
		if err != nil {
			_, rollbackErr := u.db.Rollback(ctx)
			if rollbackErr != nil {
				kgsotel.Error(ctx, rollbackErr.Error())
				err = rollbackErr
			}
			return
		}

		// Commit the transaction
		_, commitErr := u.db.Commit(ctx)
		if commitErr != nil {
			kgsotel.Error(ctx, commitErr.Error())
			err = commitErr
		}
	}()

	// Convert user status
	userStatus, kgsErr := enum.UserStatusFromInt(int(req.Status))
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	userInfo := vo.UserInfo{
		Id:       req.Id,
		Account:  req.Account,
		Password: req.Password,
		Status:   userStatus,
	}

	// Update user
	_, kgsErr = u.userService.UpdateUser(ctx, userInfo)
	if kgsErr != nil {
		return nil, kgsErr
	}

	return &auth.Empty{}, nil
}
