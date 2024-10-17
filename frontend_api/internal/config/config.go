package config

import (
	"hype-casino-platform/pkg/cfgloader"
	"log"
	"sync"
)

type (
	Host struct {
		ServiceName           string   `env:"SERVICE_NAME"`
		ServiceUrl            string   `env:"SERVICE_URL"`
		ServiceDomains        []string `env:"SERVICE_DOMAINS"`
		RateLimitIntervalSecs int      `env:"RATE_LIMIT_INTERVAL_SECS"`
		RateLimitMaxRequests  int      `env:"RATE_LIMIT_MAX_REQUESTS"`
	}

	Oauth struct {
		OauthUrl string `env:"OAUTH_URL"`
	}

	Otel struct {
		OtelUrl string `env:"OTEL_URL"`
	}

	Redis struct {
		RedisUrl    string `env:"REDIS_URL"`
		Password    string `env:"REDIS_PASSWORD"`
		DB          int    `env:"REDIS_DB"`
		MaxActive   int    `env:"REDIS_MAX_ACTIVE_CONNS"`
		MinIdle     int    `env:"REDIS_MIX_IDLE_CONNS"`
		MaxIdle     int    `env:"REDIS_MAX_IDLE_CONNS"`
		ConnTimeout int    `env:"REDIS_CONN_TIMEOUT_SECS"`
	}
	Config struct {
		Host
		Otel
		Oauth
		Redis
	}
)

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config, err := cfgloader.LoadConfigFromEnv[Config]()
		if err != nil {
			log.Fatalf("load config from env failed: %v", err)
		}
		instance = config
	})
	return instance
}
