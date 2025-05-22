package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) SetupWS() {
	s.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)

			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	s.Get("/ws/projects/:projectID/sprints/:sprintID", websocket.New(s.StreamSprint))
	s.Get("/ws/docs/:docID", websocket.New(s.StreamDocument))

	// go func() {
	// 	ticker := time.NewTicker(5 * time.Second)
	// 	defer ticker.Stop()

	// 	for range ticker.C {
	// 		if s.Logic.ProjectClient() == nil {
	// 			log.Infow("project service unavailable yet")
	// 			continue
	// 		}
	// 		// var wg sync.WaitGroup
	// 		for hubID, hub := range s.wshubs {
	// 			parts := strings.SplitN(hubID, "-", 2)
	// 			if len(parts) != 2 {
	// 				log.Warnw("invalid hub ID", "id", hubID)
	// 				continue
	// 			}
	// 			ent, id := parts[0], parts[1]
	// 			if _, err := uuid.Parse(id); err != nil {
	// 				continue
	// 			}
	// 			log.Infow("ws hub info", "type", ent, "id", id)

	// 			switch ent {
	// 			case "sprint":
	// 				func() (err error) {
	// 					task, err := s.Logic.TaskQueue.Rpop(context.Background(), hubID)
	// 					if err != nil || task.Value == nil {
	// 						log.Infow("no tasks found in queue")
	// 						return
	// 					}
	// 					if err := s.Logic.UpdateTask(context.Background(), task.Value.Id, task.Value); err != nil {
	// 						log.Errorw("failed to update task", "err", err)
	// 						return err
	// 					}
	// 					log.Infow("task is updated")

	// 					if hub.CountClient() == 0 {
	// 						return nil
	// 					}

	// 					tasks, err := s.Logic.ListTasks(context.Background(), &dto.TaskFilter{
	// 						SprintId: id,
	// 						// AssigneeId: c.Query("assignee_id"),
	// 						Page:    1,
	// 						PerPage: 10000,
	// 					})
	// 					if err != nil {
	// 						log.Errorw("failed to fetch tasks", "err", err)
	// 						return err
	// 					}
	// 					msg, err := json.Marshal(tasks.Items)
	// 					if err != nil {
	// 						log.Errorw("failed marshaling trask list", "err", err)
	// 						return
	// 					}
	// 					log.Info("broadcasting to all connected clients")

	// 					hub.Broadcast(msg)
	// 					if len(hub.GetClients()) == 0 {
	// 						log.Infow("no more clients. removing from queue", "hub_id", id)
	// 						delete(s.wshubs, id)
	// 					}
	// 					return nil
	// 				}()
	// 			case "doc":
	// 				func() (err error) {
	// 					docBody, err := s.Logic.DocumentsCache.Get(context.Background(), hubID)
	// 					if err != nil {
	// 						return err
	// 					}
	// 				}()
	// 			}
	// 		}
	// 	}
	// }()

}
