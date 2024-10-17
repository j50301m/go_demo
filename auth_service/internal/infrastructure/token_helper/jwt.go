package token_helper

import (
	"context"
	"fmt"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct{}

func NewJwtToken() *JwtToken {
	return &JwtToken{}
}

var _ TokenHelper = (*JwtToken)(nil)

func (j *JwtToken) Create(ctx context.Context, secret string, claims map[string]interface{}) (string, *kgserr.KgsError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		kgserr := kgserr.New(kgserr.InternalServerError, "Failed to create token", err)
		kgsotel.Error(ctx, kgserr.Message())
		return "", kgserr
	}
	return tokenString, nil
}

func (j *JwtToken) Validate(ctx context.Context, tokenString string, secret string) (map[string]interface{}, *kgserr.KgsError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			kgsErr := kgserr.New(kgserr.InternalServerError, fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
			kgsotel.Error(ctx, kgsErr.Message())
			return nil, kgsErr
		}

		return []byte(secret), nil
	})

	if err != nil {
		kgsErr := kgserr.New(kgserr.Unauthorized, "Jwt token is invalid", err)
		kgsotel.Error(ctx, kgsErr.Message())
		return nil, kgsErr
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	kgsErr := kgserr.New(kgserr.Unauthorized, "Invalid token claims")
	kgsotel.Error(ctx, kgsErr.Message())

	return nil, kgsErr
}

func (j *JwtToken) GetPayload(ctx context.Context, tokenString string) (map[string]interface{}, *kgserr.KgsError) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		kgsErr := kgserr.New(kgserr.Unauthorized, "Jwt token is invalid", err)
		kgsotel.Error(ctx, kgsErr.Message())
		return nil, kgsErr
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	kgsErr := kgserr.New(kgserr.Unauthorized, "Invalid token claims")
	kgsotel.Error(ctx, kgsErr.Message())

	return nil, kgsErr
}
