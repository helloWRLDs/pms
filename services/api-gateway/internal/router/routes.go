package router

import "github.com/gofiber/fiber/v2"

func (s *Server) SetupRoutes() {
	api := s.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/healthcheck", s.HealthcheckHandler)

	v1.Route("/auth", func(r fiber.Router) {
		r.Post("/login", s.LoginUser)
	})
}
