package logic

import (
	"context"
	"sync"

	"go.uber.org/zap"
	projectclient "pms.analytics/internal/clients/project"
	"pms.analytics/internal/config"
	"pms.analytics/internal/data"
	"pms.pkg/tools/scheduler"
)

type Logic struct {
	projectClient *projectclient.ProjectClient

	Conf *config.Config
	log  *zap.SugaredLogger

	Tasks map[string]*scheduler.Task

	Repo *data.Repository

	mu sync.Mutex
}

func New(repo *data.Repository, conf *config.Config, log *zap.SugaredLogger) *Logic {
	l := &Logic{
		Conf:  conf,
		log:   log,
		Tasks: make(map[string]*scheduler.Task),
		Repo:  repo,
	}

	l.InitTasks()
	for _, task := range l.Tasks {
		scheduler.Run(context.Background(), task)
	}
	return l
}
