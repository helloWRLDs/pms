package router

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/type/list"
)

func (s *Server) ListCompanies(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "ListCompanies"),
		zap.String("ip", c.IP()),
	)
	log.Debug("ListCompanies called")

	page, err := strconv.Atoi(c.Query("page", ""))
	if err != nil {
		page = 1
	}
	perPage, err := strconv.Atoi(c.Query("per_page", ""))
	if err != nil {
		perPage = 10
	}

	companies, err := s.Logic.ListCompanies(c.UserContext(), list.Filters{
		Pagination: list.Pagination{
			Page:    page,
			PerPage: perPage,
		},
	})
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
