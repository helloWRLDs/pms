package router

import "github.com/gofiber/fiber/v2"

func (s *Server) SetupRoutes() {
	api := s.Group("/api")

	s.Use(s.SecureHeaders())

	v1 := api.Group("/v1")

	v1.Get("/healthcheck", s.HealthcheckHandler)
	v1.Route("/auth", func(auth fiber.Router) {
		auth.Post("/login", s.LoginUser)
		auth.Post("/register", s.RegisterUser)
	})
	v1.Route("/background-tasks", func(tasks fiber.Router) {
		tasks.Get("/", s.ListBackgroundTasks)
	})
}
