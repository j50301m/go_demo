package httpserver

import (
	"context"
	"fmt"
	"hype-casino-platform/frontend_api/internal/config"
	"hype-casino-platform/frontend_api/internal/middleware/security"
	"hype-casino-platform/frontend_api/internal/route"
	"hype-casino-platform/pkg/kgsotel"
	otelgin "hype-casino-platform/pkg/kgsotel/gin"
	"hype-casino-platform/pkg/rate_limiter"
	"hype-casino-platform/pkg/responder"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

// NewHttpServer creates a new http server
func NewHttpServer(
	lc fx.Lifecycle,
	route route.Route,
	redisClient *redis.Client,
) *http.Server {
	// Get config
	cfg := config.GetConfig()

	// New gin server
	srv := http.Server{
		Addr: cfg.ServiceUrl,
	}

	var shutdown func(context.Context) error
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Init kgsotel
			_shutdown, err := kgsotel.InitTelemetry(ctx, cfg.Host.ServiceName, cfg.OtelUrl)
			if err != nil {
				return err
			}
			shutdown = _shutdown

			r := gin.New()
			// Register the middleware
			r.Use(otelgin.TracingMiddleware(cfg.ServiceName))
			r.Use(responder.GinResponser())
			r.Use(rate_limiter.RateLimitMiddleware(cfg.Host.ServiceName, redisClient,
				rate_limiter.WithInterval(time.Duration(cfg.Host.RateLimitIntervalSecs)*time.Second),
				rate_limiter.WithMaxRequests(int64(cfg.Host.RateLimitMaxRequests)),
			))
			// r.Use(auth.AuthMiddleware(authClient))
			r.Use(security.NewCORSMiddleware(cfg))

			// Register the routes
			route.RegisterRoutes(r)

			// Replace the handler
			srv.Handler = r

			// Listen the port
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					kgsotel.Error(ctx, "failed to serve", kgsotel.NewField("error", err))
				}
			}()

			kgsotel.Info(ctx, fmt.Sprintf("http server started at %s", cfg.ServiceUrl))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := shutdown(ctx)
			if err != nil {
				kgsotel.Error(ctx, "Error shutting down http server", kgsotel.NewField("error", err))
				return err
			}
			kgsotel.Info(ctx, "http server shut down gracefully")
			return nil
		},
	})

	return &srv
}
