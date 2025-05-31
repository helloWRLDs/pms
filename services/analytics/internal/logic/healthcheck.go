package logic

import (
	"context"
	"time"

	"go.uber.org/zap"
	authclient "pms.analytics/internal/clients/auth"
	projectclient "pms.analytics/internal/clients/project"
	"pms.pkg/tools/scheduler"
)

func (l *Logic) InitTasks() {
	healthTask := &scheduler.Task{
		ID:          "health_check",
		MaxAttempts: -1,
		Func: func(ctx context.Context) error {
			log := l.log.Named("health_check")
			log.Debug("health check task started")

			if err := l.CheckAuthHealth(ctx); err != nil {
				log.Errorw("failed to check auth health", "err", err)
			}

			if err := l.CheckProjectHealth(ctx); err != nil {
				log.Errorw("failed to check project health", "err", err)
			}

			return nil
		},
		Interval: 30 * time.Second,
	}
	l.Tasks[healthTask.ID] = healthTask
}

func (l *Logic) CheckAuthHealth(ctx context.Context) error {
	log := new(zap.SugaredLogger)
	{
		if l.Conf.Auth.DisableLog {
			log = zap.NewNop().Sugar()
		} else {
			log = l.log.With("func", "CheckAuthHealth")
		}
	}

	l.mu.Lock()
	currentAuthClient := l.authClient
	l.mu.Unlock()

	if currentAuthClient != nil {
		if currentAuthClient.Ping() {
			log.Debug("auth conn is ok")
			return nil
		}
		log.Debug("auth conn has been lost. Reconnecting...")
	}

	newClient, err := authclient.New(l.Conf.Auth, l.log)
	if err != nil {
		log.Errorw("failed to establish auth conn", "err", err)
		return err
	}

	l.mu.Lock()
	l.authClient = newClient
	l.mu.Unlock()

	log.Debug("auth conn re-established")
	return nil
}

func (l *Logic) CheckProjectHealth(ctx context.Context) error {
	log := new(zap.SugaredLogger)
	{
		if l.Conf.Project.DisableLog {
			log = zap.NewNop().Sugar()
		} else {
			log = l.log.With("func", "CheckProjectHealth")
		}
	}

	l.mu.Lock()
	currentProjectClient := l.projectClient
	l.mu.Unlock()

	if currentProjectClient != nil {
		if currentProjectClient.Ping() {
			log.Debug("project conn is ok")
			return nil
		}
		log.Debug("project conn has been lost. Reconnecting...")
	}

	newClient, err := projectclient.New(l.Conf.Project, l.log)
	if err != nil {
		log.Errorw("failed to establish project conn", "err", err)
		return err
	}

	l.mu.Lock()
	l.projectClient = newClient
	l.mu.Unlock()

	log.Debug("project conn re-established")
	return nil
}

func (l *Logic) CloseClients(ctx context.Context) {
	log := l.log.With("func", "CloseClients")

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.authClient != nil {
		if err := l.authClient.Close(); err != nil {
			log.Errorw("failed to close auth client conn", "err", err)
		} else {
			log.Debug("auth client connection closed")
		}
	}

	if l.projectClient != nil {
		if err := l.projectClient.Close(); err != nil {
			log.Errorw("failed to close project client conn", "err", err)
		} else {
			log.Debug("project client connection closed")
		}
	}
}
