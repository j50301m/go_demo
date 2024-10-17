package security

import (
	"hype-casino-platform/frontend_api/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCORSMiddleware(envCfg *config.Config) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = envCfg.ServiceDomains

	return cors.New(config)
}
