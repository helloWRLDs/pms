package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/logic"
	"pms.pkg/errs"
)

type Server struct {
	fiber.App
	Host string

	Logic *logic.Logic

	log *zap.SugaredLogger
}

func New(conf config.Config, logic *logic.Logic, log *zap.SugaredLogger) *Server {
	srv := Server{
		Host: conf.Host,
		log:  log,
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

	logger := logger.New(logger.Config{
		Format:        "[${ip}]:${port}(${locals:requestid}) ${status} - ${method} ${path}\n",
		TimeFormat:    "02-Jan-2006",
		DisableColors: false,
		Output:        os.Stdout,
	})

	srv.Use(logger)
	srv.Logic = logic
	return &srv
}

func (r *Server) Start() error {
	return r.Listen(r.Host)
}
