package config

import "pms.pkg/logger"

type Config struct {
	Host string        `env:"HOST"`
	Log  logger.Config `envPrefix:"LOG_"`
}
