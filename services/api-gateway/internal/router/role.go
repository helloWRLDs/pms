package router

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) CreateRole(c *fiber.Ctx) error {
	log := s.log.Named("CreateRole").With(
		zap.String("IP", c.IP()),
		zap.String("Method", c.Method()),
		zap.String("Path", c.Path()),
	)
	log.Info("CreateRole called")

	role := new(dto.NewRole)
	if err := c.BodyParser(role); err != nil {
		log.Error("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := s.Logic.CreateRole(c.UserContext(), role); err != nil {
		log.Error("failed to create role", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(201)
}

func (s *Server) ListRoles(c *fiber.Ctx) error {
	log := s.log.Named("ListRoles").With(
		zap.String("IP", c.IP()),
		zap.String("Method", c.Method()),
		zap.String("Path", c.Path()),
	)
	log.Info("ListRoles called")

	filter := &dto.RoleFilter{
		CompanyId:   c.Query("company_id"),
		WithDefault: c.Query("with_default", "true") == "true",
		Page:        int32(c.QueryInt("page", 1)),
		PerPage:     int32(c.QueryInt("per_page", 10)),
	}
	log.Info("ListRoles filter", zap.Any("filter", filter))

	roles, err := s.Logic.ListRoles(c.UserContext(), filter)
	if err != nil {
		log.Error("failed to list roles", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(roles)
}

func (s *Server) GetRole(c *fiber.Ctx) error {
	log := s.log.Named("GetRole").With(
		zap.String("IP", c.IP()),
		zap.String("Method", c.Method()),
		zap.String("Path", c.Path()),
	)
	log.Info("GetRole called")

	roleName := c.Params("roleName")
	if roleName == "" {
		log.Error("role name is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "role name is required",
		})
	}

	role, err := s.Logic.GetRole(c.UserContext(), roleName)
	if err != nil {
		log.Error("failed to get role", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(role)
}

func (s *Server) UpdateRole(c *fiber.Ctx) error {
	log := s.log.Named("UpdateRole").With(
		zap.String("IP", c.IP()),
		zap.String("Method", c.Method()),
		zap.String("Path", c.Path()),
	)
	log.Info("UpdateRole called")

	roleName := c.Params("roleName")
	if roleName == "" {
		log.Error("role name is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "role name is required",
		})
	}

	role := new(dto.Role)
	if err := c.BodyParser(role); err != nil {
		log.Error("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if c.Locals("company_id") == nil {
		log.Error("company_id is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "company_id is required",
		})
	}
	companyID := c.Locals("company_id").(string)
	if role.CompanyId != companyID {
		log.Error("role company_id does not match")
		return errs.ErrForbidden{
			Reason: "cannot update role for another company",
		}
	}

	if err := s.Logic.UpdateRole(c.UserContext(), roleName, role, companyID); err != nil {
		log.Error("failed to update role", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(200)
}

func (s *Server) DeleteRole(c *fiber.Ctx) error {
	log := s.log.Named("DeleteRole").With(
		zap.String("IP", c.IP()),
		zap.String("Method", c.Method()),
		zap.String("Path", c.Path()),
	)
	log.Info("DeleteRole called")

	roleName := c.Params("roleName")
	if roleName == "" {
		log.Error("role name is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "role name is required",
		})
	}

	if c.Locals("company_id") == nil {
		log.Error("company_id is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "company_id is required",
		})
	}
	companyID := c.Locals("company_id").(string)

	if err := s.Logic.DeleteRole(c.UserContext(), roleName, companyID); err != nil {
		log.Error("failed to delete role", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(200)
}
