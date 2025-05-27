package logic

import (
	"context"
	"time"

	"go.uber.org/zap"
	analyticsclient "pms.api-gateway/internal/client/analytics"
	authclient "pms.api-gateway/internal/client/auth"
	notifierclient "pms.api-gateway/internal/client/notifier"
	projectclient "pms.api-gateway/internal/client/project"
	"pms.pkg/tools/scheduler"
)

func (l *Logic) InitTasks() {
	analyticsTask := &scheduler.Task{
		ID:          "analytics-client-connector",
		MaxAttempts: -1,
		Func:        l.CheckAnalyticsHealth,
		Interval:    10 * time.Second,
	}
	authTask := &scheduler.Task{
		ID:          "auth-client-connector",
		MaxAttempts: -1,
		Func:        l.CheckAuthHealth,
		Interval:    10 * time.Second,
	}
	notifierTask := &scheduler.Task{
		ID:          "notifier-mq-connector",
		MaxAttempts: -1,
		Func:        l.CheckNotifierHealth,
		Interval:    10 * time.Second,
	}
	projectTask := &scheduler.Task{
		ID:          "project-client-connector",
		MaxAttempts: -1,
		Func:        l.CheckProjectHealth,
		Interval:    10 * time.Second,
	}

	l.Tasks[authTask.ID] = authTask
	l.Tasks[notifierTask.ID] = notifierTask
	l.Tasks[projectTask.ID] = projectTask
	l.Tasks[analyticsTask.ID] = analyticsTask
}

func (l *Logic) CheckAnalyticsHealth(ctx context.Context) error {
	log := new(zap.SugaredLogger)
	{
		if l.Config.Analytics.DisableLog {
			log = zap.NewNop().Sugar()
		} else {
			log = l.log.With("func", "CheckAnalyticsHealth")
		}
	}

	l.mu.Lock()
	currentAnalyticsClient := l.analyticsClient
	l.mu.Unlock()

	if currentAnalyticsClient != nil {
		if currentAnalyticsClient.Ping() {
			log.Debug("analytics conn is ok")
			return nil
		}
		log.Debug("analytics conn has been lost. Reconnecting...")
	}

	newClient, err := analyticsclient.New(l.Config.Analytics, l.log)
	if err != nil {
		log.Errorw("failed to establish project conn", "err", err)
		return err
	}

	l.mu.Lock()
	l.analyticsClient = newClient
	l.mu.Unlock()

	log.Info("analytics conn re-established")
	return nil

}

func (l *Logic) CheckProjectHealth(ctx context.Context) error {
	log := new(zap.SugaredLogger)
	{
		if l.Config.Project.DisableLog {
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

	newClient, err := projectclient.New(l.Config.Project, l.log)
	if err != nil {
		log.Errorw("failed to establish project conn", "err", err)
		return err
	}

	l.mu.Lock()
	l.projectClient = newClient
	l.mu.Unlock()

	log.Info("project conn re-established")
	return nil

}

func (l *Logic) CheckAuthHealth(ctx context.Context) error {
	log := new(zap.SugaredLogger)
	{
		if l.Config.Auth.DisableLog {
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

	newClient, err := authclient.New(l.Config.Auth, l.log)
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

func (l *Logic) CheckNotifierHealth(ctx context.Context) error {
	log := new(zap.SugaredLogger)
	{
		if l.Config.NotificationMQ.DisableLog {
			log = zap.NewNop().Sugar()
		} else {
			log = l.log.With("func", "CheckNotifierHealth")
		}
	}
	log.Debug("CheckNotifierHealth called")

	l.mu.Lock()
	currentMQClient := l.notificationMQ // Copy to avoid accessing while unlocked
	l.mu.Unlock()

	if currentMQClient != nil {
		if currentMQClient.ConnState() == "READY" {
			log.Debug("notifier conn is ok")
			return nil
		}
		// log.Debug("notifier conn has been lost. Reconnecting...")
	}

	newClient, err := notifierclient.New(l.Config.NotificationMQ, l.log)
	if err != nil {
		// log.Errorw("failed to establish notifier conn", "err", err)
		return err
	}

	l.mu.Lock()
	l.notificationMQ = newClient
	l.mu.Unlock()

	log.Info("notifier conn re-established")
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

	if l.notificationMQ != nil {
		if err := l.notificationMQ.Close(); err != nil {
			log.Errorw("failed to close notifier client conn", "err", err)
		} else {
			log.Debug("notifier client connection closed")
		}
	}
	if l.projectClient != nil {
		l.projectClient.Close()
	}

	if l.analyticsClient != nil {
		l.analyticsClient.Close()
	}
}
