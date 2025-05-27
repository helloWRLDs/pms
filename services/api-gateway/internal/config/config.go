package config

import (
	"pms.pkg/datastore/mq"
	"pms.pkg/datastore/redis"
	"pms.pkg/logger"
	"pms.pkg/tools/jwtoken"
	configgrpc "pms.pkg/transport/grpc/config"
)

type Config struct {
	Host        string `env:"HOST"`
	FrontendURL string `env:"FRONTEND_URL"`

	JWT jwtoken.Config `envPrefix:"JWT_"`

	Redis redis.Config `envPrefix:"REDIS_"`

	Auth           configgrpc.ClientConfig `envPrefix:"AUTH_"`
	Project        configgrpc.ClientConfig `envPrefix:"PROJECT_"`
	Analytics      configgrpc.ClientConfig `envPrefix:"ANALYTICS_"`
	NotificationMQ mq.Config               `envPrefix:"NOTIFICATION_"`
	// Notifier notifierclient.Config `envPrefix:"NOTIFIER_"`

	Log logger.Config `envPrefix:"LOG_"`
}
