package config

import (
	authclient "pms.api-gateway/internal/client/auth"
	notifierclient "pms.api-gateway/internal/client/notifier"
	"pms.pkg/logger"
)

type Config struct {
	Host     string                `env:"HOST"`
	Log      logger.Config         `envPrefix:"LOG_"`
	Auth     authclient.Config     `envPrefix:"AUTH_"`
	Notifier notifierclient.Config `envPrefix:"NOTIFIER_"`
}
