package aggregate

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgserr"
)

type User struct {
	Id                    int64                                                             // Primary key for the user
	Account               string                                                            // Unique name for each user
	Password              string                                                            // Hashed password
	PasswordFailTimes     int                                                               // Number of times the user has failed to login
	Status                enum.UserStatus                                                   // Status of the user
	clientLoader          func(ctx context.Context) (*Client, *kgserr.KgsError)             // Lazy loader for the client
	roleLoader            func(ctx context.Context) (*entity.Role, *kgserr.KgsError)        // Lazy loader for the role
	lastLoginRecordLoader func(ctx context.Context) (*entity.LoginRecord, *kgserr.KgsError) // Lazy loader for the last login record
}

func (u *User) SetRoleLoader(loader func(ctx context.Context) (*entity.Role, *kgserr.KgsError)) {
	u.roleLoader = loader
}

func (u *User) Role(ctx context.Context) (*entity.Role, *kgserr.KgsError) {
	return u.roleLoader(ctx)
}

func (u *User) SetLoginRecordLoader(loader func(ctx context.Context) (*entity.LoginRecord, *kgserr.KgsError)) {
	u.lastLoginRecordLoader = loader
}

func (u *User) LoginRecord(ctx context.Context) (*entity.LoginRecord, *kgserr.KgsError) {
	return u.lastLoginRecordLoader(ctx)
}

func (u *User) SetClientLoader(loader func(ctx context.Context) (*Client, *kgserr.KgsError)) {
	u.clientLoader = loader
}

func (u *User) Client(ctx context.Context) (*Client, *kgserr.KgsError) {
	return u.clientLoader(ctx)
}
