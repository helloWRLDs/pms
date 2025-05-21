package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) ListTaskComments(c *fiber.Ctx) error {
	log := s.log.Named("ListTaskComments").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("ListTaskComments called")

	filter := &dto.TaskCommentFilter{
		Page:    int32(c.QueryInt("page", 1)),
		PerPage: int32(c.QueryInt("per_page", 10)),
		TaskId:  c.Params("taskID", ""),
		UserId:  c.Query("user_id", ""),
	}

	commentList, err := s.Logic.ListTaskComments(c.UserContext(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(commentList)
}

func (s *Server) CreateTaskComments(c *fiber.Ctx) error {
	log := s.log.Named("CreateTaskComments").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("CreateTaskComments called")

	creation := new(dto.TaskCommentCreation)
	if err := c.BodyParser(creation); err != nil {
		return errs.ErrBadGateway{
			Object: "task comment",
		}
	}

	created, err := s.Logic.CreateTaskComment(c.UserContext(), creation)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(created)
}

func (s *Server) GetTaskComment(c *fiber.Ctx) error {
	log := s.log.Named("GetTaskComment").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("GetTaskComment called")

	commentID := c.Params("commentID", "")
	if s := strings.Trim(commentID, " "); s == "" {
		return errs.ErrBadGateway{
			Object: "task-comment-id",
		}
	}

	comment, err := s.Logic.GetTaskComment(c.UserContext(), commentID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(comment)
}
