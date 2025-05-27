package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
			EnablePrintRoutes: false,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				if err == nil {
					return nil
				}

				http := errs.WrapHttp(err).(errs.ErrHTTP)
				return c.Status(http.Status).JSON(http)
			},
		}),
	}

	srv.Use(cors.New())
	srv.Use(requestid.New())
	srv.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: time.Minute,
	}))

	srv.Logic = logic
	return &srv
}

func (r *Server) Start() error {
	return r.Listen(r.Host)
}
