package config

import (
	analyticsclient "pms.api-gateway/internal/client/analytics"
	authclient "pms.api-gateway/internal/client/auth"
	projectclient "pms.api-gateway/internal/client/project"
	"pms.pkg/datastore/mq"
	"pms.pkg/datastore/redis"
	"pms.pkg/logger"
	"pms.pkg/tools/jwtoken"
)

type Config struct {
	Host string `env:"HOST"`

	JWT jwtoken.Config `envPrefix:"JWT_"`

	Redis redis.Config `envPrefix:"REDIS_"`

	Auth           authclient.Config      `envPrefix:"AUTH_"`
	Project        projectclient.Config   `envPrefix:"PROJECT_"`
	NotificationMQ mq.Config              `envPrefix:"NOTIFICATION_"`
	Analytics      analyticsclient.Config `envPrefix:"ANALYTICS_"`
	// Notifier notifierclient.Config `envPrefix:"NOTIFIER_"`

	Log logger.Config `envPrefix:"LOG_"`
}
