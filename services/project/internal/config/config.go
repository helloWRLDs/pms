package config

import (
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
)

type Config struct {
	Host string          `env:"HOST"`
	DB   postgres.Config `envPrefix:"POSTGRES_"`
	Log  logger.Config   `envPrefix:"LOG_"`
}
