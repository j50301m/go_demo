package repository

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/aggregate"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/pkg/kgserr"
)

type UserRepo interface {
	Find(ctx context.Context, id int64) (*aggregate.User, *kgserr.KgsError)
	Create(ctx context.Context, clientId int64, user *aggregate.User) (*aggregate.User, *kgserr.KgsError)
	Update(ctx context.Context, user *aggregate.User) (*aggregate.User, *kgserr.KgsError)
	AddLoginRecord(ctx context.Context, userId int64, loginRecord *entity.LoginRecord) (*entity.LoginRecord, *kgserr.KgsError)
	BindRole(ctx context.Context, userId int64, roleId int64) (*aggregate.User, *kgserr.KgsError)
}
