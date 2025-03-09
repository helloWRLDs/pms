package router

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Server) HealthcheckHandler(c *fiber.Ctx) error {
	type ServiceState struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}

	res := struct {
		Services []ServiceState `json:"services"`
	}{}

	if auth := r.Logic.AuthClient(); auth != nil {
		res.Services = append(res.Services, ServiceState{
			Name:   "auth-service",
			Status: auth.State().String(),
		})
	} else {
		res.Services = append(res.Services, ServiceState{
			Name:   "auth-service",
			Status: "DOWN",
		})
	}
	if notifier := r.Logic.NotificationQueue(); notifier != nil {
		res.Services = append(res.Services, ServiceState{
			Name:   "notifier-mq",
			Status: notifier.ConnState(),
		})
	} else {
		res.Services = append(res.Services, ServiceState{
			Name:   "notifier-mq",
			Status: "DOWN",
		})
	}

	return c.Status(200).JSON(res)
}
