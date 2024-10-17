package main

import (
	"hype-casino-platform/frontend_api/internal/infrastructure/grpc_client"
	httpserver "hype-casino-platform/frontend_api/internal/infrastructure/http_server"
	"hype-casino-platform/frontend_api/internal/infrastructure/redis_initializer"
	"hype-casino-platform/frontend_api/internal/route"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		grpc_client.NewGrpcClientSet(),
		route.NewRouteV1Set(),
		fx.Provide(
			httpserver.NewHttpServer,
			redis_initializer.NewRedisClient,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
