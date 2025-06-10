package router

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) CreateTask(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "CreateTask"),
		zap.String("ip", c.IP()),
	)
	log.Debug("CreateTask called")

	creation := new(dto.TaskCreation)
	if err := c.BodyParser(creation); err != nil {
		log.Errorw("failed to parse task creation", "err", err)
		return err
	}

	log.Infow("task creation", "creation", creation)

	if err := s.Logic.CreateTask(c.UserContext(), creation); err != nil {
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

	log.Infow("check locals", "project_id", c.Locals("project_id"))

	filter := &dto.TaskFilter{
		Page:        int32(c.QueryInt("page", 1)),
		PerPage:     int32(c.QueryInt("per_page", 10)),
		SprintId:    c.Query("sprint_id", ""),
		SprintName:  c.Query("sprint_name", ""),
		ProjectId:   c.Locals("project_id").(string),
		ProjectName: c.Query("project_name", ""),
		AssigneeId:  c.Query("assignee_id", ""),
		Status:      c.Query("status", ""),
		Priority:    int32(c.QueryInt("priority", 0)),
		Type:        c.Query("type", ""),
	}
	log.Infow("filters", "filters", filter)

	tasks, err := s.Logic.ListTasks(c.UserContext(), filter)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(tasks)
}

func (s *Server) UpdateTask(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "UpdateTask"),
		zap.String("ip", c.IP()),
	)
	log.Debug("UpdateTask called")

	taskID := c.Params("taskID", "")
	if strings.Trim(taskID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "task_id",
		}
	}

	toUpdate := new(dto.Task)
	if err := c.BodyParser(toUpdate); err != nil {
		log.Errorw("failed to parse task update", "err", err)
		return err
	}
	log.Infow("parsed task body", "task", toUpdate)

	if err := s.Logic.UpdateTask(c.UserContext(), taskID, toUpdate); err != nil {
		log.Errorw("failed to update task", "err", err)
		return err
	}
	return c.SendStatus(http.StatusOK)
}

func (s *Server) DeleteTask(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "DeleteTask"),
		zap.String("ip", c.IP()),
	)
	log.Debug("DeleteTask called")

	taskID := c.Params("taskID", "")
	if strings.Trim(taskID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "task_id",
		}
	}

	log.Infow("deleting task", "task_id", taskID)

	if err := s.Logic.DeleteTask(c.UserContext(), taskID); err != nil {
		log.Errorw("failed to delete task", "err", err)
		return err
	}

	return c.SendStatus(200)
}

func (s *Server) CreateTaskAssignment(c *fiber.Ctx) error {
	log := s.log.Named("CreateTaskAssignment").With(
		zap.String("ip", c.IP()),
	)
	log.Info("CreateTaskAssignment called")

	taskID := c.Params("taskID", "")
	if strings.Trim(taskID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "task_id",
		}
	}
	userID := c.Params("userID", "")
	if strings.Trim(userID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "user_id",
		}
	}
	log.Infow("creating task assignment", "task_id", taskID, "user_id", userID)

	if err := s.Logic.TaskAssign(c.UserContext(), taskID, userID); err != nil {
		return err
	}

	return c.SendStatus(200)

}

func (s *Server) DeleteTaskAssignment(c *fiber.Ctx) error {
	log := s.log.Named("DeleteTaskAssignment").With(
		zap.String("ip", c.IP()),
	)
	log.Info("DeleteTaskAssignment called")

	taskID := c.Params("taskID", "")
	if strings.Trim(taskID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "task_id",
		}
	}
	userID := c.Params("userID", "")
	if strings.Trim(userID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "user_id",
		}
	}

	log.Infow("deleting task assignment", "task_id", taskID, "user_id", userID)

	if err := s.Logic.TaskUnassign(c.UserContext(), taskID, userID); err != nil {
		return err
	}

	return c.SendStatus(200)
}
