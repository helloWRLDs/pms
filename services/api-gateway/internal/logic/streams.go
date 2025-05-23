package logic

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/transport/ws"
)

func (l *Logic) processDocumentStream() {
	log := l.log.Named("processDocumentStream")

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		if l.ProjectClient() == nil || l.AnalyticsClient() == nil {
			log.Debugw("project or analytics service unavailable yet")
			continue
		}

		var wg sync.WaitGroup
		for hubID, hub := range l.WsHubs {
			wg.Add(1)
			go func(hubID string, hub *ws.Hub) {
				defer wg.Done()

				parts := strings.SplitN(hubID, "-", 2)
				if len(parts) != 2 {
					log.Warnw("invalid hub ID", "id", hubID)
					return
				}
				name, id := parts[0], parts[1]
				if _, err := uuid.Parse(id); err != nil {
					return
				}
				if _, err := uuid.Parse(id); err != nil {
					return
				}
				log.Debugw("ws hub info", "type", name, "id", id)
				if !strings.HasPrefix(hubID, "doc") {
					return
				}
				cachedDoc, err := l.DocumentsCache.Get(context.Background(), hubID)
				if err != nil {
					return
				}
				raw, err := json.Marshal(cachedDoc.Document)
				if err != nil {
					log.Errorw("failed marshaling doc body", "err", err)
					return
				}
				hub.Broadcast(raw)

				if cachedDoc.RequireUpdate {
					updateRes, err := l.analyticsClient.UpdateDocument(context.Background(), &pb.UpdateDocumentRequest{
						DocId:      cachedDoc.Document.Id,
						UpdatedDoc: cachedDoc.Document,
					})
					log.Debugw("update doc results", "res", updateRes)
					if err != nil {
						log.Errorw("failed to update doc", "err", err)
					}
					cachedDoc.RequireUpdate = false
					l.DocumentsCache.Set(context.Background(), hubID, cachedDoc, 24)
				}
			}(hubID, hub)
		}
	}
}

func (l *Logic) processTaskStream() {
	log := l.log.Named("processTaskStream")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if l.ProjectClient() == nil {
			log.Debug("project service unavailable yet")
			continue
		}

		var wg sync.WaitGroup
		for hubID, hub := range l.WsHubs {
			wg.Add(1)
			go func(hubID string, hub *ws.Hub) {
				defer wg.Done()

				parts := strings.SplitN(hubID, "-", 2)
				if len(parts) != 2 {
					log.Warnw("invalid hub ID", "id", hubID)
					return
				}
				name, id := parts[0], parts[1]
				if _, err := uuid.Parse(id); err != nil {
					return
				}

				if !strings.HasPrefix(hubID, "sprint") {
					return
				}

				log.Debugw("ws hub info", "type", name, "id", id)

				task, err := l.TaskQueue.Rpop(context.Background(), hubID)
				if err != nil || task.Value == nil {
					log.Debugw("no tasks found in queue")
					return
				}
				if err := l.UpdateTask(context.Background(), task.Value.Id, task.Value); err != nil {
					log.Errorw("failed to update task", "err", err)
					return
				}
				log.Debug("task is updated")

				if hub.CountClient() == 0 {
					log.Infow("no clients found")
					return
				}

				tasks, err := l.ListTasks(context.Background(), &dto.TaskFilter{
					SprintId: id,
					Page:     1,
					PerPage:  10000,
				})
				if err != nil {
					log.Errorw("failed to fetch tasks", "err", err)
					return
				}
				msg, err := json.Marshal(tasks.Items)
				if err != nil {
					log.Errorw("failed marshaling trask list", "err", err)
					return
				}
				log.Debug("broadcasting to all connected clients")

				hub.Broadcast(msg)
				if len(hub.GetClients()) == 0 {
					log.Debugw("no more clients. removing from queue", "hub_id", id)
					delete(l.WsHubs, hubID)
				}
			}(hubID, hub)

		}
		wg.Wait()
	}
}
