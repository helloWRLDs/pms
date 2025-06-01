package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) UpdateUser(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "UpdateUser"),
		zap.String("ip", c.IP()),
	)
	log.Debug("UpdateUser called")

	userID := c.Params("id", "")
	if strings.Trim(userID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "user_id",
		}
	}

	var user dto.User
	if err := c.BodyParser(&user); err != nil {
		return errs.ErrBadGateway{
			Object: "user",
		}
	}

	err := s.Logic.UpdateUser(c.UserContext(), userID, &user)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func (s *Server) GetUser(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "GetUserProfile"),
		zap.String("ip", c.IP()),
	)
	log.Debug("GetUserProfile called")

	userID := c.Params("id", "")
	if strings.Trim(userID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "user_id",
		}
	}

	user, err := s.Logic.GetUserProfile(c.UserContext(), userID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(user)
}

func (s *Server) ListUsers(c *fiber.Ctx) error {
	log := s.log.Named("ListUsers").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("ListUsers called")

	filter := &dto.UserFilter{
		Page:      int32(c.QueryInt("page", 1)),
		PerPage:   int32(c.QueryInt("per_page", 10)),
		CompanyId: c.Query("company_id", ""),
	}

	userList, err := s.Logic.ListUsers(c.UserContext(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(userList)
}
