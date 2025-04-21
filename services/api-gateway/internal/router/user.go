package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
)

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
