package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) CreateRole(ctx context.Context, role *dto.NewRole) error {
	log := l.log.Named("CreateRole").With(
		zap.String("name", role.Name),
		zap.String("companyID", role.CompanyId),
		zap.Any("permissions", role.Permissions),
	)
	log.Info("CreateRole called")

	roleRes, err := l.authClient.CreateRole(ctx, &pb.CreateRoleRequest{
		Role: role,
	})
	if err != nil {
		log.Error("failed to create role", zap.Error(err))
		return err
	}

	log.Infow("role created", "res", zap.Any("role", roleRes))

	return nil
}

func (l *Logic) ListRoles(ctx context.Context, filter *dto.RoleFilter) (*dto.RoleList, error) {
	log := l.log.Named("ListRoles").With(
		zap.Any("filter", filter),
	)
	log.Info("ListRoles called")

	roleRes, err := l.authClient.ListRoles(ctx, &pb.ListRolesRequest{
		Filter: filter,
	})
	if err != nil {
		log.Error("failed to list roles", zap.Error(err))
		return nil, err
	}

	log.Infow("roles listed", "res", zap.Any("roles", roleRes))

	return roleRes.Roles, nil
}

func (l *Logic) GetRole(ctx context.Context, name string) (*dto.Role, error) {
	log := l.log.Named("GetRole").With(
		zap.String("name", name),
	)
	log.Info("GetRole called")

	roleRes, err := l.authClient.GetRole(ctx, &pb.GetRoleRequest{
		Name: name,
	})
	if err != nil {
		log.Error("failed to get role", zap.Error(err))
		return nil, err
	}

	log.Infow("role retrieved", "res", zap.Any("role", roleRes))

	return roleRes.Role, nil
}

func (l *Logic) UpdateRole(ctx context.Context, name string, role *dto.Role, companyID string) error {
	log := l.log.Named("UpdateRole").With(
		zap.String("name", name),
		zap.Any("role", role),
	)
	log.Info("UpdateRole called")

	roleRes, err := l.authClient.UpdateRole(ctx, &pb.UpdateRoleRequest{
		Name:      name,
		Role:      role,
		CompanyId: companyID,
	})
	if err != nil {
		log.Error("failed to update role", zap.Error(err))
		return err
	}

	log.Infow("role updated", "res", zap.Any("role", roleRes))

	return nil
}

func (l *Logic) DeleteRole(ctx context.Context, name string, companyID string) error {
	log := l.log.Named("DeleteRole").With(
		zap.String("name", name),
		zap.String("companyID", companyID),
	)
	log.Info("DeleteRole called")

	roleRes, err := l.authClient.DeleteRole(ctx, &pb.DeleteRoleRequest{
		Name:      name,
		CompanyId: companyID,
	})
	if err != nil {
		log.Error("failed to delete role", zap.Error(err))
		return err
	}

	log.Infow("role deleted", "res", zap.Any("role", roleRes))
	return nil
}
