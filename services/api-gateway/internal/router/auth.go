package router

import (
	"github.com/gofiber/fiber/v2"
	"pms.pkg/errs"
	"pms.pkg/protobuf/dto"
)

func (s *Server) LoginUser(c *fiber.Ctx) error {
	var creds dto.UserCreds
	err := c.BodyParser(&creds)
	if err != nil {
		return errs.WrapGRPC(errs.ErrInvalidInput{
			Object: "user credentials",
			Reason: "failed to resolve credentials",
		})
	}

	return c.SendStatus(200)
}
