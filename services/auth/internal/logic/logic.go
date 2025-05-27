package logic

import (
	"go.uber.org/zap"
	"pms.auth/internal/config"
	"pms.auth/internal/data"
	"pms.pkg/api/google"
)

type Logic struct {
	Repo *data.Repository
	conf *config.Config

	log *zap.SugaredLogger

	googleClient *google.Client
}

func New(repo *data.Repository, conf *config.Config, log *zap.SugaredLogger) *Logic {
	return &Logic{
		Repo:         repo,
		conf:         conf,
		log:          log,
		googleClient: google.New(conf.GoogleConfig, log),
	}
}
