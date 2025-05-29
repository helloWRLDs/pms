package config

import (
	"pms.pkg/api/github"
	"pms.pkg/api/google"
	"pms.pkg/datastore/postgres"
	"pms.pkg/logger"
	"pms.pkg/tools/jwtoken"
)

type Config struct {
	Host         string          `env:"HOST"`
	DB           postgres.Config `envPrefix:"POSTGRES_"`
	Log          logger.Config   `envPrefix:"LOG_"`
	JWT          jwtoken.Config  `envPrefix:"JWT_"`
	GoogleConfig google.Config   `envPrefix:"GOOGLE_"`
	GitHubConfig github.Config   `envPrefix:"GITHUB_"`
}
