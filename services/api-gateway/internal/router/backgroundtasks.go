package router

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *Server) ListBackgroundTasks(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "router.ListBackgroundTasks"),
		zap.String("ip", c.IP()),
	)
	log.Debug("ListBackgroundTasks called")

	type Task struct {
		ID          string `json:"id"`
		MaxAttempts int    `json:"max_attempts,omitempty"`
		Status      string `json:"status"`
		Attempts    int    `json:"attempts,omitempty"`
	}

	tasks := make([]Task, 0)
	for _, task := range s.Logic.Tasks {
		t := Task{
			ID:     task.ID,
			Status: task.Status().String(),
		}
		if task.MaxAttempts > 0 {
			t.MaxAttempts = task.MaxAttempts
			t.Attempts = task.Attempts()
		}
		tasks = append(tasks, t)
	}
	return c.Status(200).JSON(tasks)
}
