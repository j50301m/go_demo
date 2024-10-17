package db

import (
	"context"
	"hype-casino-platform/pkg/kgserr"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, *kgserr.KgsError)
	Set(ctx context.Context, key string, value string, expiration time.Duration) *kgserr.KgsError
	GetObject(ctx context.Context, key string, dest any) *kgserr.KgsError
	SetObject(ctx context.Context, key string, value any, expiration time.Duration) *kgserr.KgsError
	Delete(ctx context.Context, keys ...string) *kgserr.KgsError
}
