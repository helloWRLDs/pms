package router

import (
	"github.com/gofiber/fiber/v2"
	"pms.pkg/errs"
	"pms.pkg/protobuf/dto"
)

func (s *Server) LoginUser(c *fiber.Ctx) error {
	var creds dto.UserCredentials
	err := c.BodyParser(&creds)
	if err != nil {
		return errs.WrapGRPC(errs.ErrInvalidInput{
			Object: "user credentials",
			Reason: "failed to resolve credentials",
		})
	}
	payload, err := s.Logic.LoginUser(c.UserContext(), &creds)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(payload)
}

func (s *Server) RegisterUser(c *fiber.Ctx) error {
	var newUser dto.NewUser
	if err := c.BodyParser(&newUser); err != nil {
		return errs.ErrInvalidInput{
			Object: "new user data",
			Reason: "failed to parse user data",
		}
	}
	created, err := s.Logic.RegisterUser(c.UserContext(), &newUser)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(created)
}
