package router

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
)

func (s *Server) GetCompanyStats(c *fiber.Ctx) error {
	log := s.log.Named("GetProjectStats").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("GetProjectStats called")

	companyID := c.Params("companyID")
	if companyID == "" {
		return errs.ErrBadGateway{
			Object: "projectID",
		}
	}

	stats, err := s.Logic.GetProjectStats(c.UserContext(), companyID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(stats)
}
