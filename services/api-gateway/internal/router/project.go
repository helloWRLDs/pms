package router

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func (s *Server) ListProjects(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "ListProjects"),
		zap.String("ip", c.IP()),
	)
	log.Info("ListProjects called")

	filter := &dto.ProjectFilter{
		Page:      utils.If(c.Query("page", "") != "", int32(c.QueryInt("page", 1)), 1),
		PerPage:   utils.If(c.Query("per_page", "") != "", int32(c.QueryInt("per_page", 10)), 10),
		CompanyId: c.Query("company_id", ""),
		Title:     c.Query("title", ""),
		Status:    c.Query("status", ""),
	}
	log.Debugw("filter", "filter", filter)

	projects, err := s.Logic.ListProjects(c.UserContext(), filter)
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

	projectID := c.Params("projectID", "")
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
