package logic

import (
	"context"
	"sync"

	"go.uber.org/zap"
	authclient "pms.analytics/internal/clients/auth"
	projectclient "pms.analytics/internal/clients/project"
	"pms.analytics/internal/config"
	"pms.analytics/internal/data"
	"pms.pkg/tools/scheduler"
)

type Logic struct {
	projectClient *projectclient.ProjectClient
	authClient    *authclient.AuthClient

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

	authClient, err := authclient.New(conf.Auth, log)
	if err != nil {
		log.Errorw("failed to create auth client", "err", err)
	} else {
		l.authClient = authClient
	}

	projectClient, err := projectclient.New(conf.Project, log)
	if err != nil {
		log.Errorw("failed to create project client", "err", err)
	} else {
		l.projectClient = projectClient
	}

	l.InitTasks()
	for _, task := range l.Tasks {
		scheduler.Run(context.Background(), task)
	}
	return l
}
