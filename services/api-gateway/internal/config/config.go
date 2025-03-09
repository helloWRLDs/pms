package config

import (
	authclient "pms.api-gateway/internal/client/auth"
	"pms.pkg/datastore/mq"
	"pms.pkg/datastore/redis"
	"pms.pkg/logger"
	"pms.pkg/tools/jwtoken"
)

type Config struct {
	Host string `env:"HOST"`

	JWT jwtoken.Config `envPrefix:"JWT_"`

	Redis redis.Config `envPrefix:"REDIS_"`

	Auth           authclient.Config `envPrefix:"AUTH_"`
	NotificationMQ mq.Config         `envPrefix:"NOTIFICATION_"`
	// Notifier notifierclient.Config `envPrefix:"NOTIFIER_"`

	Log logger.Config `envPrefix:"LOG_"`
}
