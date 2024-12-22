package router

import (
	"github.com/gofiber/fiber/v2"
	"pms.api-gateway/internal/config"
	"pms.pkg/errs"
)

type Server struct {
	fiber.App
	Host string
}

func New(conf config.Config) *Server {
	return &Server{
		Host: conf.Host,
		App: *fiber.New(fiber.Config{
			AppName:           "API-GATEWAY",
			EnablePrintRoutes: true,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				if err == nil {
					return nil
				}

				http := errs.WrapHttp(err).(errs.ErrHTTP)
				return c.Status(http.Status).JSON(http)
			},
		}),
	}
}
