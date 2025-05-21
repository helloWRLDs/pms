package logic

import (
	"context"
	"sync"

	"go.uber.org/zap"
	authclient "pms.api-gateway/internal/client/auth"
	projectclient "pms.api-gateway/internal/client/project"
	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/models"
	"pms.pkg/datastore/mq"
	"pms.pkg/datastore/redis"
	"pms.pkg/tools/scheduler"
	"pms.pkg/transport/ws"
)

type Logic struct {
	Config config.Config

	authClient    *authclient.AuthClient
	projectClient *projectclient.ProjectClient

	notificationMQ *mq.Publisher

	TaskQueue *redis.Client[models.TaskQueueElement]
	Sessions  *redis.Client[models.Session]
	Tasks     map[string]*scheduler.Task

	ws.Hub

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

func (l *Logic) ProjectClient() *projectclient.ProjectClient {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.projectClient
}

func New(config config.Config, log *zap.SugaredLogger) *Logic {
	l := &Logic{
		Config:     config,
		log:        log,
		stopTicker: make(chan struct{}),
		Tasks:      make(map[string]*scheduler.Task),
		Sessions:   redis.New(&config.Redis, models.Session{}),
		TaskQueue:  redis.New(&config.Redis, models.TaskQueueElement{}),
	}
	l.InitTasks()
	for _, task := range l.Tasks {
		scheduler.Run(context.Background(), task)
	}
	return l
}
