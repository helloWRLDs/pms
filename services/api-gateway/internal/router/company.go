package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) ListCompanies(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "ListCompanies"),
		zap.String("ip", c.IP()),
	)
	log.Debug("ListCompanies called")

	filter := &dto.CompanyFilter{
		Page:        int32(c.QueryInt("page", 1)),
		PerPage:     int32(c.QueryInt("per_page", 10)),
		UserId:      c.Query("user_id", ""),
		CompanyId:   c.Query("company_id", ""),
		CodeName:    c.Query("company_codename", ""),
		CompanyName: c.Query("company_name", ""),
	}

	companies, err := s.Logic.ListCompanies(c.UserContext(), filter)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(companies)
}

func (s *Server) GetCompany(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "GetCompany"),
		zap.String("ip", c.IP()),
	)
	log.Debug("GetCompany called")

	companyID := c.Params("id", "")
	if strings.Trim(companyID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "company_id",
		}
	}

	company, err := s.Logic.GetCompany(c.UserContext(), companyID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(company)
}
