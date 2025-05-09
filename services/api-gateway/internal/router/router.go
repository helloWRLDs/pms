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

	v1.Route("/session", func(session fiber.Router) {
		session.Use(s.Authorize())

		session.Get("/", s.GetSession)
		session.Put("/", s.UpdateSession)
		session.Delete("/", s.DeleteSession)
	})

	v1.Route("/users", func(user fiber.Router) {
		user.Use(s.RequireAuthService())
		user.Use(s.Authorize())

		user.Get("/:id", s.GetUser)
	})

	v1.Route("/companies", func(comp fiber.Router) {
		comp.Use(s.RequireAuthService())
		comp.Use(s.Authorize())

		comp.Get("/", s.ListCompanies)
		comp.Get("/:id", s.GetCompany)
	})

	v1.Route("/projects", func(proj fiber.Router) {
		proj.Use(s.RequireAuthService())
		proj.Use(s.Authorize())

		proj.Get("/", s.ListProjects)
		proj.Get("/:id", s.GetProject)
		proj.Post("/", s.CreateProject)
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
