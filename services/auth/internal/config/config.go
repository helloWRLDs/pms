package config

import (
	"pms.pkg/datastore/sqlite"
	"pms.pkg/logger"
	"pms.pkg/tools/jwt"
)

type Config struct {
	Host string        `env:"HOST"`
	DB   sqlite.Config `envPrefix:"SQLITE_"`
	Log  logger.Config `envPrefix:"LOG_"`
	JWT  jwt.Config    `envPrefix:"JWT_"`
}
