package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/jwtoken"
	"pms.pkg/type/claims"
	"pms.pkg/utils"
	ctxutils "pms.pkg/utils/ctx"
)

func (s *Server) CheckCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := s.Logic.GetSessionInfo(c.UserContext())
		if err != nil {
			return err
		}

		project := c.Params("projectID", "")

		if strings.Trim(project, " ") == "" {
			return errs.ErrBadGateway{
				Object: "project_id",
			}
		}

		if utils.ContainsInArray(session.Projects, c.Params("projectID", "")) {
			return errs.ErrUnauthorized{
				Reason: "don't have access to project",
			}
		}
		log.Info("project_id: ", project)
		c.Locals("project_id", project)

		return c.Next()
	}
}

func (s *Server) Authorize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log := s.log.With(
			zap.String("func", "Authorize"),
		)
		log.Debug("Authorize called")

		token := c.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			log.Error("missing token")
			return errs.ErrUnauthorized{
				Reason: "token not provided",
			}
		}
		token = token[7:]

		log.Debugw("jwt secrets", "ttl", s.Logic.Config.JWT.TTL, "secret", s.Logic.Config.JWT.Secret)
		decodedRaw, err := jwtoken.DecodeToken(token, &claims.AccessTokenClaims{}, &s.Logic.Config.JWT)
		if err != nil {
			log.Errorw("failed decoding token", "err", err)
			return errs.ErrUnauthorized{
				Reason: "failed verifying token",
			}
		}

		decoded, ok := decodedRaw.(*claims.AccessTokenClaims)
		if !ok {
			log.Error("failed to cast token claims")
			return errs.ErrUnauthorized{
				Reason: "invalid token claims",
			}
		}

		log.Debugf("claims: %#v", decoded)
		session, err := s.Logic.Sessions.Get(c.UserContext(), decoded.SessionID)
		if err != nil {
			log.Error("failed to get session from cache")
			return errs.ErrUnauthorized{
				Reason: "failed verifying session",
			}
		}
		log.Debugw("got session from cache", "session", session)
		ctx := ctxutils.Embed(c.UserContext(), session)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

func (s *Server) RequireAuthService() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if s.Logic.AuthClient() == nil {
			return errs.ErrUnavalaiable{
				Object: "auth Service",
			}
		}
		return c.Next()
	}
}

func (s *Server) SecureHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Frame-Options", "deny")
		c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		return c.Next()
	}
}
