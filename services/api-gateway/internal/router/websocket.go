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

	s.Get("/ws/docs/:docID", websocket.New(s.StreamDocument))
	s.Get("/ws/sprints/:sprintID", websocket.New(s.StreamSprint))
}
