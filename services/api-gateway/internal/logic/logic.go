package logic

import (
	"context"
	"sync"

	"go.uber.org/zap"
	authclient "pms.api-gateway/internal/client/auth"
	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/models"
	"pms.pkg/datastore/mq"
	"pms.pkg/datastore/redis"
	"pms.pkg/tools/scheduler"
)

type Logic struct {
	Config config.Config

	authClient     *authclient.AuthClient
	notificationMQ *mq.Publisher

	Sessions *redis.Client[models.Session]
	Tasks    map[string]*scheduler.Task

	stopTicker chan struct{}
	log        *zap.SugaredLogger

	mu sync.Mutex
}

func (l *Logic) NotificationQueue() *mq.Publisher {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.notificationMQ
}

func (l *Logic) AuthClient() *authclient.AuthClient {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.authClient
}

func New(config config.Config, log *zap.SugaredLogger) *Logic {
	l := &Logic{
		Config:     config,
		log:        log,
		stopTicker: make(chan struct{}),
		Tasks:      make(map[string]*scheduler.Task),
		Sessions:   redis.New(&config.Redis, models.Session{}),
	}
	l.InitTasks()
	for _, task := range l.Tasks {
		scheduler.Run(context.Background(), task)
	}
	return l
}
