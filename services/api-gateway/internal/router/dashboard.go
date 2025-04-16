package router

import (
	"encoding/json"
	"time"

	"github.com/gofiber/contrib/websocket"
	"pms.pkg/transport/ws"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Executor  string    `json:"executor"`
	Status    string    `json:"status"`
	Backlog   string    `json:"backlog_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	tasks = []Task{
		{ID: "1", Title: "Setup project", Body: "Initialize repo and project setup", Status: "CREATED", Executor: "test", Backlog: "1"},
		{ID: "2", Title: "Implement login", Body: "Create login page and API", Status: "DONE", Executor: "test", Backlog: "1"},
		{ID: "3", Title: "Design dashboard", Body: "Dashboard wireframe and design", Status: "PENDING", Executor: "test", Backlog: "1"},
	}
)

func (s *Server) DashboardStream(c *websocket.Conn) {
	log := s.log.With("func", "DashboardStream")
	log.Debug("DashboardStream called")

	s.DashboardHub.AddClient(c)
	defer s.DashboardHub.RemoveClient(c)

	sendTasks := func() {
		data, err := json.Marshal(tasks)
		if err != nil {
			log.Errorw("failed to marshal tasks", "err", err)
			return
		}
		s.DashboardHub.Broadcast(data)
	}
	sendTasks()

	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			// log.Errorw("failed to read message", "err", err)
			break
		}

		log.Infow("received msg", "msg", string(msg), "mt", mt)

		var payload ws.Action[Task]
		if err := json.Unmarshal(msg, &payload); err == nil && payload.Action == "update" {
			log.Infow("received payload", "payload", payload)
			for i, task := range tasks {
				if task.ID == payload.ID {
					tasks[i] = payload.Value
					s.DashboardHub.SetCache(payload.ID, payload.Value)
					break
				}
			}
			// Broadcast updated task list to all clients
			sendTasks()
		} else {
			log.Errorw("failed to unmarshal payload", "err", err, "payload", string(msg))
		}
	}
}
