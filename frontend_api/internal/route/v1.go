package route

import (
	v1_handler "hype-casino-platform/frontend_api/internal/api/v1"
	"hype-casino-platform/frontend_api/internal/middleware/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// RouteV1 is a struct implementing the Route interface
// Version 1 of the API
type RouteV1 struct {
	authClient  auth.AuthClient // This client is used to validate token for the middleware
	authHandler *v1_handler.AuthHandler
	// Add other handlers here
}

var _ Route = (*RouteV1)(nil)

// NewRouteV1Set creates a new fx.Option for the RouteV1 module
func NewRouteV1Set() fx.Option {
	return fx.Module("route-v1",
		fx.Provide(
			newRouteV1,
			v1_handler.NewAuthHandler,
			// Add other handlers here
		),
	)
}

// newRouteV1 creates a new RouteV1 instance for the Route interface
func newRouteV1(
	authClient auth.AuthClient,
	authHandler *v1_handler.AuthHandler) Route {
	return &RouteV1{
		authClient:  authClient,
		authHandler: authHandler,
	}
}

func (r *RouteV1) RegisterRoutes(g *gin.Engine) {
	v1 := g.Group("/api/v1")
	g.Use(auth.AuthMiddleware(
		r.authClient,
		auth.ByPassPath("/api/v1/auth/*")),
	)
	r.addAuthRoutes(v1)
}

func (r *RouteV1) addAuthRoutes(g *gin.RouterGroup) {
	auth := g.Group("/auth")
	auth.GET("/", r.authHandler.ClientAuth)
	auth.POST("/login", r.authHandler.Login)
}
