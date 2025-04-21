package logic

import (
	"go.uber.org/zap"
	"pms.auth/internal/config"
	"pms.auth/internal/data"
)

type Logic struct {
	Repo *data.Repository
	conf *config.Config

	log *zap.SugaredLogger
}

func New(repo *data.Repository, conf *config.Config, log *zap.SugaredLogger) *Logic {
	return &Logic{
		Repo: repo,
		conf: conf,
		log:  log,
	}
}
