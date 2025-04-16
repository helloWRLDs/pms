package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/logic"
	"pms.pkg/errs"
	"pms.pkg/transport/ws"
)

type Server struct {
	fiber.App
	Host string

	Logic *logic.Logic

	DashboardHub *ws.Hub

	log *zap.SugaredLogger
}

func New(conf config.Config, logic *logic.Logic, log *zap.SugaredLogger) *Server {
	srv := Server{
		DashboardHub: ws.NewHub(),
		Host:         conf.Host,
		log:          log,
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

	srv.Use(requestid.New())
	srv.Use(cors.New())

	srv.Logic = logic
	return &srv
}

func (r *Server) Start() error {
	return r.Listen(r.Host)
}
