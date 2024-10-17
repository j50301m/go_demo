package grpc_client

import (
	"hype-casino-platform/frontend_api/internal/middleware/auth"

	"go.uber.org/fx"
)

func NewGrpcClientSet() fx.Option {
	return fx.Module("grpc-client",
		fx.Provide(
			NewAuthClient,
			fx.Annotate(
				func(a *AuthClient) auth.AuthClient { return a },
			),
			// Add other clients here
		),
	)
}
