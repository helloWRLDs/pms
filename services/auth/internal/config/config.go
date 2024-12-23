package config

import (
	"pms.pkg/datastore/sqlite"
	"pms.pkg/logger"
)

type Config struct {
	Host string        `env:"HOST"`
	DB   sqlite.Config `envPrefix:"SQLITE_"`
	Log  logger.Config `envPrefix:"LOG_"`
}
