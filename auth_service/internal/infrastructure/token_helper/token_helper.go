package token_helper

import (
	"context"
	"hype-casino-platform/pkg/kgserr"
)

type TokenHelper interface {
	Create(ctx context.Context, secret string, claims map[string]any) (string, *kgserr.KgsError)
	Validate(ctx context.Context, token string, secret string) (map[string]any, *kgserr.KgsError)
	GetPayload(ctx context.Context, token string) (map[string]any, *kgserr.KgsError)
}
