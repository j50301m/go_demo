package grpc_impl

import (
	"context"
	"fmt"
	"hype-casino-platform/auth_service/internal/application"
	"hype-casino-platform/auth_service/internal/config"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
	otelgrpc "hype-casino-platform/pkg/kgsotel/grpc"
	"hype-casino-platform/pkg/pb/gen/auth"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewGrpcServer(lc fx.Lifecycle,
	authService *application.AuthService,
	clientService *application.ClientService,
	userService *application.UserService) *grpc.Server {
	// Get config
	cfg := config.GetConfig()

	// New grpc server
	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.TracingMiddleware(otelgrpc.RoleServer)),
		grpc.ChainUnaryInterceptor(
			kgserr.ErrorInterceptor,
			// Any other interceptors can be added here
		),
		grpc.StreamInterceptor(
			kgserr.StreamErrorInterceptor,
			// Any other interceptors can be added here
		),
	)

	var shutdown func(context.Context) error
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Init kgsotel
			_shutdown, err := kgsotel.InitTelemetry(ctx, cfg.Host.ServiceName, cfg.OtelUrl)
			if err != nil {
				return err
			}
			shutdown = _shutdown

			// Listen the port
			lis, err := net.Listen("tcp", cfg.ServiceUrl)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			// Register the service
			go func() {
				auth.RegisterAuthServiceServer(s, authService)
				auth.RegisterClientServiceServer(s, clientService)
				auth.RegisterUserServiceServer(s, userService)
				if err := s.Serve(lis); err != nil {
					kgsotel.Error(ctx, "failed to serve", kgsotel.NewField("error", err))
				}
			}()

			kgsotel.Info(ctx, fmt.Sprintf("gRPC server started at %s", cfg.ServiceUrl))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := shutdown(ctx)
			if err != nil {
				return err
			}

			kgsotel.Info(ctx, "gRPC server shut down gracefully")
			return nil
		},
	})

	return s
}
