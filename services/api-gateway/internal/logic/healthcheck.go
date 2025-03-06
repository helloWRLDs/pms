package logic

import (
	"context"
	"time"

	authclient "pms.api-gateway/internal/client/auth"
	"pms.pkg/tools/scheduler"
)

func (l *Logic) InitTasks() {
	authTask := &scheduler.Task{
		ID:          "auth-client-connector",
		MaxAttempts: -1,
		Func:        l.CheckAuthHealth,
		Interval:    30 * time.Second,
	}
	l.Tasks[authTask.ID] = authTask
}

func (l *Logic) CheckAuthHealth(ctx context.Context) error {
	log := l.log.With("func", "CheckAuthHealth")
	if l.AuthClient != nil {
		if l.AuthClient.Ping() {
			log.Debug("auth conn is ok")
			return nil
		}
		log.Debug("auth conn has been lost. Reconnecting...")
		l.AuthClient = nil
	}

	newClient, err := authclient.New(l.Config.Auth, l.log)
	if err != nil {
		log.Errorw("failed to establish conn", "err", err)
		return err
	}
	l.AuthClient = newClient
	log.Info("conn established")
	return nil
}

func (l *Logic) CloseClients(ctx context.Context) {
	log := l.log.With("func", "CloseClients")

	if err := l.AuthClient.Close(); err != nil {
		log.Errorw("failed to close auth client conn", "err", err)
	}
}
