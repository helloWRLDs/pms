package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/utils/ctx"
)

// Find session in redis and put into context
func (s *Server) Authorize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log := s.log.With(
			zap.String("func", "Authorize"),
		)
		log.Debug("Authorize called")

		token := c.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			return errs.ErrUnauthorized{
				Reason: "token not provided",
			}
		}
		token = token[7:]

		decoded, err := s.Logic.Config.JWT.DecodeToken(token)
		if err != nil {
			return errs.ErrUnauthorized{
				Reason: "failed verifying token",
			}
		}
		session, err := s.Logic.Sessions.Get(c.UserContext(), decoded.SessionID)
		if err != nil {
			return errs.ErrUnauthorized{
				Reason: "failed verifying session",
			}
		}
		c.SetUserContext(ctx.Embed(c.Context(), session))
		return c.Next()
	}
}

func (s *Server) SecureHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Frame-Options", "deny")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		return c.Next()
	}
}
