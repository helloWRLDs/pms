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

func (s *Server) CreateTask(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "CreateTask"),
		zap.String("ip", c.IP()),
	)
	log.Debug("CreateTask called")

	var creation dto.TaskCreation
	if err := c.BodyParser(&creation); err != nil {
		log.Errorw("failed to parse task creation", "err", err)
		return err
	}
	log.Infow("task creation", "creation", creation)

	if err := s.Logic.CreateTask(c.UserContext(), &creation); err != nil {
		log.Errorw("failed to create task", "err", err)
		return err
	}

	return c.SendStatus(http.StatusCreated)
}

func (s *Server) GetTask(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "GetTask"),
		zap.String("ip", c.IP()),
	)
	log.Debug("GetTask called")

	taskID := c.Params("taskID", "")
	if strings.Trim(taskID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "task_id",
		}
	}

	task, err := s.Logic.GetTask(c.UserContext(), taskID)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(task)
}

func (s *Server) ListTasks(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "ListTasks"),
		zap.String("ip", c.IP()),
	)
	log.Debug("ListTasks called")

	filters := list.Filters{
		Pagination: list.Pagination{
			Page:    c.QueryInt("page", 1),
			PerPage: c.QueryInt("per_page", 10),
		},
		Fields: map[string]string{
			"sprint_id":   c.Query("sprint_id", ""),
			"assignee_id": c.Query("assignee_id", ""),
			"project_id":  c.Params("projectID", ""),
		},
	}

	tasks, err := s.Logic.ListTasks(c.UserContext(), filters)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(tasks)
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
