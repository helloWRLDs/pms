package router

import (
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (s *Server) SetupREST() {
	api := s.Group("/api")

	s.Use(s.SecureHeaders())

	v1 := api.Group("/v1")

	v1.Get("/healthcheck", s.HealthcheckHandler)

	v1.Route("/auth", func(auth fiber.Router) {
		auth.Use(s.RequireAuthService())

		auth.Post("/login", s.LoginUser)
		auth.Post("/register", s.RegisterUser)
	})

	v1.Route("/background-tasks", func(tasks fiber.Router) {
		tasks.Get("/", s.ListBackgroundTasks)
	})
}

func (s *Server) SetupWS() {
	s.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)

			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	s.Get("/ws/dashboard/:id", websocket.New(s.DashboardStream))

	go func() { // dashboard chages db writer
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if len(s.DashboardHub.GetClients()) == 0 {
				log.Infow("no clients connected")
				continue
			}
			if len(s.DashboardHub.GetCache()) == 0 {
				log.Infow("no tasks to save")
				continue
			}
			log.Infow("saving tasks to db...", "tasks", s.DashboardHub.GetCache())
			s.DashboardHub.Clean()
		}
	}()
}
