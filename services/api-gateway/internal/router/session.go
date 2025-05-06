package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.api-gateway/internal/models"
)

func (s *Server) GetSession(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "GetSession"),
		zap.String("ip", c.IP()),
	)
	log.Debug("GetSession called")

	session, err := s.Logic.GetSessionInfo(c.UserContext())
	if err != nil {
		return err
	}

	return c.Status(200).JSON(session)
}

func (s *Server) UpdateSession(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "UpdateSession"),
		zap.String("ip", c.IP()),
	)
	log.Debug("UpdateSession called")

	session, err := s.Logic.GetSessionInfo(c.UserContext())
	if err != nil {
		return err
	}

	var updatedSession models.Session
	if err := c.BodyParser(&updatedSession); err != nil {
		return err
	}
	updatedSession.ID = session.ID

	err = s.Logic.Sessions.Set(c.UserContext(), updatedSession.ID, updatedSession, int64(session.Expires.Hour()))
	if err != nil {
		log.Errorw("failed to update session", "err", err)
		return err
	}
	log.Debugw("session updated", "session", updatedSession)

	return c.Status(200).JSON(fiber.Map{
		"msg": fmt.Sprintf("session %s updated", updatedSession.ID),
	})
}

func (s *Server) DeleteSession(c *fiber.Ctx) error {
	log := s.log.With(
		zap.String("func", "DeleteSession"),
		zap.String("ip", c.IP()),
	)
	log.Debug("DeleteSession called")

	session, err := s.Logic.GetSessionInfo(c.UserContext())
	if err != nil {
		return err
	}

	if err := s.Logic.Sessions.Delete(c.UserContext(), session.ID); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"msg": fmt.Sprintf("session %s deleted", session.ID),
	})
}
