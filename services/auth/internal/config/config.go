package config

import (
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
	"pms.pkg/tools/jwtoken"
)

type Config struct {
	Host string `env:"HOST"`
	// DB   sqlite.Config  `envPrefix:"SQLITE_"`
	DB  postgres.Config `envPrefix:"POSTGRES_"`
	Log logger.Config   `envPrefix:"LOG_"`
	JWT jwtoken.Config  `envPrefix:"JWT_"`
}
