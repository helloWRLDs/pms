package router

import "github.com/gofiber/fiber/v2"

func (r *Server) SetupRoutes() {
	api := r.Group("/api")
	api.Route("/v1", func(v1 fiber.Router) {
		v1.Get("/healthcheck", r.HealthcheckHandler)
	})
}

func (r *Server) Start() error {
	return r.Listen(r.Host)
}
