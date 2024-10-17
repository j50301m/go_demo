package repository

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/aggregate"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/pkg/kgserr"
)

type ClientRepo interface {
	Create(ctx context.Context, client *aggregate.Client) (*aggregate.Client, *kgserr.KgsError)
	Find(ctx context.Context, id int64) (*aggregate.Client, *kgserr.KgsError)
	Update(ctx context.Context, client *aggregate.Client) (*aggregate.Client, *kgserr.KgsError)
	BindSystemRoles(ctx context.Context, clientId int64, sysRoles ...entity.Role) *kgserr.KgsError
	CreateRoles(ctx context.Context, clientId int64, roles ...entity.Role) ([]entity.Role, *kgserr.KgsError)
	DeleteRoles(ctx context.Context, clientId int64, roleIds ...int64) *kgserr.KgsError
	UpdateRoles(ctx context.Context, clientId int64, roles ...entity.Role) ([]entity.Role, *kgserr.KgsError)
	FindRole(ctx context.Context, clientId int64, roleId int64) (*entity.Role, *kgserr.KgsError)
}
