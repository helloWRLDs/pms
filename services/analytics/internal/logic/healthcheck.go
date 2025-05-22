package logic

import (
	"context"
	"time"

	projectclient "pms.analytics/internal/clients/project"
	"pms.pkg/tools/scheduler"
)

func (l *Logic) InitTasks() {
	projectTask := &scheduler.Task{
		ID:          "project-client-connector",
		MaxAttempts: -1,
		Func:        l.CheckProjectHealth,
		Interval:    10 * time.Second,
	}

	l.Tasks[projectTask.ID] = projectTask
}

func (l *Logic) CheckProjectHealth(ctx context.Context) error {
	log := l.log.With("func", "CheckProjectHealth")

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

	if l.projectClient != nil {
		if err := l.projectClient.Close(); err != nil {
			log.Errorw("failed to close project client conn", "err", err)
		} else {
			log.Debug("project client connection closed")
		}
	}

}
