package config

import (
	projectclient "pms.analytics/internal/clients/project"
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
)

type Config struct {
	Host    string               `env:"HOST"`
	DB      postgres.Config      `envPrefix:"POSTGRES_"`
	Log     logger.Config        `envPrefix:"LOG_"`
	Project projectclient.Config `envPrefix:"PROJECT_"`
}
