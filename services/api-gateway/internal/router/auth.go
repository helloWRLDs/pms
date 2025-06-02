package router

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) OAuth2Callback(c *fiber.Ctx) error {
	provider := c.Params("provider")
	if provider == "" {
		return errs.ErrInvalidInput{
			Object: "provider",
			Reason: "provider is required",
		}
	}
	code := c.Query("code")
	if code == "" {
		return errs.ErrInvalidInput{
			Object: "code",
			Reason: "code is required",
		}
	}

	user, payload, err := s.Logic.CompleteOAuth2(c.UserContext(), provider, code)
	if err != nil {

		errorURL := fmt.Sprintf("%s/auth/callback?error=%s",
			s.Logic.Config.FrontendURL,
			url.QueryEscape(err.Error()))
		return c.Redirect(errorURL)
	}

	log.Infow("user logged in", "user", user, "payload", payload)

	authResponse := map[string]interface{}{
		"user":    user,
		"payload": payload,
	}

	authJSON, err := json.Marshal(authResponse)
	if err != nil {
		log.Errorw("failed to marshal auth response", "error", err)
		errorURL := fmt.Sprintf("%s/auth/callback?error=internal_error", s.Logic.Config.FrontendURL)
		return c.Redirect(errorURL)
	}

	authData := base64.URLEncoding.EncodeToString(authJSON)

	callbackURL := fmt.Sprintf("%s/auth/callback?success=true&data=%s",
		s.Logic.Config.FrontendURL,
		url.QueryEscape(authData))

	return c.Redirect(callbackURL)
}

func (s *Server) InitiateOAuth2(c *fiber.Ctx) error {
	provider := c.Params("provider")
	if provider == "" {
		return errs.ErrInvalidInput{
			Object: "provider",
			Reason: "provider is required",
		}
	}
	authURL, err := s.Logic.InitiateOAuth2(c.UserContext(), provider)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"auth_url": authURL,
	})
}

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
