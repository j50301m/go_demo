package main

import (
	"hype-casino-platform/auth_service/internal/application"
	"hype-casino-platform/auth_service/internal/domain/repository"
	"hype-casino-platform/auth_service/internal/domain/service"
	"hype-casino-platform/auth_service/internal/infrastructure/db_impl"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl"
	"hype-casino-platform/auth_service/internal/infrastructure/grpc_impl"
	"hype-casino-platform/auth_service/internal/infrastructure/redis_impl"
	"hype-casino-platform/auth_service/internal/infrastructure/token_helper"
	"hype-casino-platform/pkg/db"
	redis_cache "hype-casino-platform/pkg/db/redis"
	"hype-casino-platform/pkg/req_analyzer"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	fx.New(
		db_impl.NewEntDbFx(),
		fx.Provide(
			grpc_impl.NewGrpcServer,
			application.NewAuthService,
			application.NewClientService,
			application.NewUserService,
			service.NewAuthService,
			service.NewClientService,
			service.NewUserService,
			fx.Annotate(
				ent_impl.NewClientRepoImpl,
				fx.As(new(repository.ClientRepo)),
			),
			fx.Annotate(
				ent_impl.NewUserRepoImpl,
				fx.As(new(repository.UserRepo)),
			),
			fx.Annotate(
				token_helper.NewJwtToken,
				fx.As(new(token_helper.TokenHelper)),
			),
			fx.Annotate(
				redis_cache.NewRedisCache,
				fx.As(new(db.Cache)),
			),
			redis_impl.NewRedisClient,
			req_analyzer.NewReqAnalyzer,
		),
		fx.Invoke(func(server *grpc.Server) {}),
	).Run()
}
