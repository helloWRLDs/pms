package config

import (
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
)

type Config struct {
	Host    string                  `env:"HOST"`
	DB      postgres.Config         `envPrefix:"POSTGRES_"`
	Log     logger.Config           `envPrefix:"LOG_"`
	Auth    configgrpc.ClientConfig `envPrefix:"AUTH_"`
	Project configgrpc.ClientConfig `envPrefix:"PROJECT_"`
}
