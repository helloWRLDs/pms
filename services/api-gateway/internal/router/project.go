package router

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
)

func (s *Server) ListProjects(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "ListProjects"),
		zap.String("ip", c.IP()),
	)
	log.Debug("ListProjects called")

	company_id := c.Get("company_id", "")
	if strings.Trim(company_id, " ") == "" {
		return errs.ErrBadGateway{
			Object: "company_id",
		}
	}

	projects, err := s.Logic.ListProjects(c.UserContext(), company_id, list.Filters{Pagination: list.Pagination{Page: 1, PerPage: 10}})
	if err != nil {
		return err
	}

	return c.Status(200).JSON(projects)
}

func (s *Server) GetProject(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "GetProject"),
		zap.String("ip", c.IP()),
	)
	log.Debug("GetProject called")

	projectID := c.Params("id", "")
	if strings.Trim(projectID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "project_id",
		}
	}

	project, err := s.Logic.GetProjectByID(c.UserContext(), projectID)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(project)
}

func (s *Server) CreateProject(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "CreateProject"),
		zap.String("ip", c.IP()),
	)
	log.Debug("CreateProject called")

	var creation dto.ProjectCreation
	if err := c.BodyParser(&creation); err != nil {
		return err
	}

	if err := s.Logic.CreateProject(c.UserContext(), &creation); err != nil {
		return err
	}

	return c.SendStatus(http.StatusCreated)
}
