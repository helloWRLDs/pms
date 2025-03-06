package logic

import (
	"go.uber.org/zap"
	"pms.auth/internal/config"
	"pms.auth/internal/data"
)

type Logic struct {
	repo *data.Repository
	conf *config.Config

	log *zap.SugaredLogger
}

func New(repo *data.Repository, conf *config.Config, log *zap.SugaredLogger) *Logic {
	return &Logic{
		repo: repo,
		conf: conf,
		log:  log,
	}
}
