package logic

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	roledata "pms.auth/internal/data/role"

	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func (l *Logic) CreateRole(ctx context.Context, newRole *dto.NewRole) error {
	log := l.log.With(
		zap.String("func", "CreateRole"),
		zap.Any("new_role", newRole),
	)
	log.Debug("CreateRole called")

	role := roledata.Role{
		Name:        newRole.Name,
		CompanyID:   utils.Ptr(newRole.CompanyId),
		Permissions: newRole.Permissions,
		CreatedAt:   time.Now(),
	}
	if err := l.Repo.Role.Create(ctx, role); err != nil {
		log.Errorw("failed to create role", "err", err)
	}
	return nil
}

func (l *Logic) GetRole(ctx context.Context, name string) (*dto.Role, error) {
	log := l.log.With(
		zap.String("func", "GetRole"),
		zap.String("name", name),
	)
	log.Debug("GetRole called")

	role, err := l.Repo.Role.GetByName(ctx, name)
	if err != nil {
		log.Errorw("failed to get role", "err", err)
	}
	return role.DTO(), nil
}

func (l *Logic) ListRoles(ctx context.Context, filter *dto.RoleFilter) (result *dto.RoleList, err error) {
	log := l.log.With(
		zap.String("func", "ListRoles"),
		zap.Any("filter", filter),
	)
	log.Debug("ListRoles called")

	roles, err := l.Repo.Role.List(ctx, filter)
	if err != nil {
		log.Errorw("failed to list roles", "err", err)
		return nil, err
	}
	result = new(dto.RoleList)
	result.Page = int32(roles.Page)
	result.PerPage = int32(roles.PerPage)
	result.TotalPages = int32(roles.TotalPages)
	result.TotalItems = int32(roles.TotalItems)
	result.Items = make([]*dto.Role, len(roles.Items))
	for i, role := range roles.Items {
		result.Items[i] = role.DTO()
	}

	return result, nil
}

func (l *Logic) UpdateRole(ctx context.Context, name string, updatedRole *dto.Role, companyID string) error {
	log := l.log.With(
		zap.String("func", "UpdateRole"),
		zap.String("name", name),
		zap.Any("updated_role", updatedRole),
	)
	log.Debug("UpdateRole called")

	existing, err := l.Repo.Role.GetByName(ctx, name)
	if err != nil {
		log.Errorw("failed to get role", "err", err)
		return err
	}

	if existing.CompanyID == nil {
		return errs.ErrForbidden{
			Reason: "cannot update default role",
		}
	}
	if *existing.CompanyID != companyID {
		return errs.ErrForbidden{
			Reason: fmt.Sprintf("role %s is not associated with company %s", name, companyID),
		}
	}

	existing.Name = updatedRole.Name
	existing.CompanyID = &updatedRole.CompanyId
	existing.Permissions = updatedRole.Permissions

	if err := l.Repo.Role.Update(ctx, name, existing); err != nil {
		log.Errorw("failed to update role", "err", err)
		return err
	}
	return nil
}

func (l *Logic) DeleteRole(ctx context.Context, name string, companyID string) error {
	log := l.log.With(
		zap.String("func", "DeleteRole"),
		zap.String("name", name),
	)
	log.Debug("DeleteRole called")

	existing, err := l.Repo.Role.GetByName(ctx, name)
	if err != nil {
		log.Errorw("failed to get role", "err", err)
		return err
	}

	if existing.CompanyID == nil {
		return errs.ErrForbidden{
			Reason: "cannot delete default role",
		}
	}
	if *existing.CompanyID != companyID {
		return errs.ErrForbidden{
			Reason: fmt.Sprintf("role %s is not associated with company %s", name, companyID),
		}
	}

	if err := l.Repo.Role.Delete(ctx, name); err != nil {
		log.Errorw("failed to delete role", "err", err)
		return err
	}
	return nil
}
