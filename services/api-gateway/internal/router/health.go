package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r *Server) HealthcheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
