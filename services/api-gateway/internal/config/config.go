package config

import (
	authclient "pms.api-gateway/internal/client/auth"
	notifierclient "pms.api-gateway/internal/client/notifier"
	"pms.api-gateway/internal/modules/cache"
	"pms.pkg/logger"
	"pms.pkg/tools/jwt"
)

type Config struct {
	Host string `env:"HOST"`

	JWT jwt.Config `envPrefix:"JWT_"`

	Redis cache.Config `envPrefix:"REDIS_"`

	Auth     authclient.Config     `envPrefix:"AUTH_"`
	Notifier notifierclient.Config `envPrefix:"NOTIFIER_"`

	Log logger.Config `envPrefix:"LOG_"`
}
