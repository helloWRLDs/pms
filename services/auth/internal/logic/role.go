package logic

import (
	"context"

	"go.uber.org/zap"
	roledata "pms.auth/internal/data/role"

	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
)

func (l *Logic) CreateRole(ctx context.Context, newRole *dto.NewRole) error {
	log := l.log.With(
		zap.String("func", "CreateRole"),
		zap.Any("new_role", newRole),
	)
	log.Debug("CreateRole called")

	role := roledata.Role{
		Name:        newRole.Name,
		CompanyID:   &newRole.CompanyId,
		Permissions: newRole.Permissions,
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

func (l *Logic) ListRoles(ctx context.Context, filter *dto.RoleFilter) (result list.List[*dto.Role], err error) {
	log := l.log.With(
		zap.String("func", "ListRoles"),
		zap.Any("filter", filter),
	)
	log.Debug("ListRoles called")

	roleFilter := dto.RoleFilter{
		Page:           filter.Page,
		PerPage:        filter.PerPage,
		DateFrom:       filter.DateFrom,
		DateTo:         filter.DateTo,
		OrderBy:        filter.OrderBy,
		OrderDirection: filter.OrderDirection,
		CompanyId:      filter.CompanyId,
		CompanyName:    filter.CompanyName,
		Name:           filter.Name,
	}

	roles, err := l.Repo.Role.List(ctx, &roleFilter)
	if err != nil {
		log.Errorw("failed to list roles", err)
		return list.List[*dto.Role]{}, err
	}
	result.Page = roles.Page
	result.PerPage = roles.PerPage
	result.TotalPages = roles.TotalPages
	result.TotalItems = roles.TotalItems

	return result, nil
}

func (l *Logic) UpdateRole(ctx context.Context, name string, updatedRole *dto.NewRole) error {
	log := l.log.With(
		zap.String("func", "UpdateRole"),
		zap.String("name", name),
		zap.Any("updated_role", updatedRole),
	)
	log.Debug("UpdateRole called")

	role := roledata.Role{
		Name:        updatedRole.Name,
		CompanyID:   &updatedRole.CompanyId,
		Permissions: updatedRole.Permissions,
	}
	if err := l.Repo.Role.Update(ctx, name, role); err != nil {
		log.Errorw("failed to update role", "err", err)
		return err
	}
	return nil
}

func (l *Logic) DeleteRole(ctx context.Context, name string) error {
	log := l.log.With(
		zap.String("func", "DeleteRole"),
		zap.String("name", name),
	)
	log.Debug("DeleteRole called")

	if err := l.Repo.Role.Delete(ctx, name); err != nil {
		log.Errorw("failed to delete role", "err", err)
		return err
	}
	return nil
}
