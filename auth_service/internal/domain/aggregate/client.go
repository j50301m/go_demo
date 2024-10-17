package aggregate

import (
	"context"

	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgserr"
)

type Client struct {
	Id               int64
	MerchantId       int64
	ClientType       enum.Client
	Secret           string
	Active           bool
	TokenExpireSecs  int
	LoginFailedTimes int
	rolesLoader      func(ctx context.Context) (*map[int64]entity.Role, *kgserr.KgsError)
}

func (c *Client) Roles(ctx context.Context) (*map[int64]entity.Role, *kgserr.KgsError) {
	return c.rolesLoader(ctx)
}

func (c *Client) SetRolesLoader(loader func(ctx context.Context) (*map[int64]entity.Role, *kgserr.KgsError)) {
	c.rolesLoader = loader
}
