package grpchandler

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	log := s.log.With(
		zap.Any("request", req),
	)
	log.Debug("CreateRole called")

	if err := s.logic.CreateRole(ctx, req.Role); err != nil {
		log.Errorw("failed to create role", "err", err)
		return nil, status.Error(codes.Internal, "failed to create role")
	}

	// Get the created role to return in response
	role, err := s.logic.GetRole(ctx, req.Role.Name)
	if err != nil {
		log.Errorw("failed to get created role", "err", err)
		return nil, status.Error(codes.Internal, "failed to get created role")
	}

	return &pb.CreateRoleResponse{
		Success: true,
		Role:    role,
	}, nil
}

func (s *ServerGRPC) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	log := s.log.With(
		zap.String("name", req.Name),
		zap.String("company_id", req.CompanyId),
	)
	log.Debug("GetRole called")

	role, err := s.logic.GetRole(ctx, req.Name)
	if err != nil {
		log.Errorw("failed to get role", "err", err)
		return nil, status.Error(codes.Internal, "failed to get role")
	}

	return &pb.GetRoleResponse{
		Success: true,
		Role:    role,
	}, nil
}

func (s *ServerGRPC) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	log := s.log.With(
		zap.String("company_id", req.CompanyId),
	)
	log.Debug("ListRoles called")

	filter := &dto.RoleFilter{
		CompanyId: req.CompanyId,
	}
	roles, err := s.logic.ListRoles(ctx, filter)
	if err != nil {
		log.Errorw("failed to list roles", "err", err)
		return nil, status.Error(codes.Internal, "failed to list roles")
	}

	return &pb.ListRolesResponse{
		Success: true,
		Roles:   roles.Items,
	}, nil
}

func (s *ServerGRPC) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	log := s.log.With(
		zap.String("name", req.Name),
		zap.String("company_id", req.CompanyId),
		zap.Any("role", req.Role),
	)
	log.Debug("UpdateRole called")

	// Convert dto.Role to dto.NewRole
	newRole := &dto.NewRole{
		Name:        req.Role.Name,
		Permissions: req.Role.Permissions,
		CompanyId:   req.Role.CompanyId,
	}

	if err := s.logic.UpdateRole(ctx, req.Name, newRole); err != nil {
		log.Errorw("failed to update role", "err", err)
		return nil, status.Error(codes.Internal, "failed to update role")
	}

	// Get the updated role to return in response
	role, err := s.logic.GetRole(ctx, req.Role.Name)
	if err != nil {
		log.Errorw("failed to get updated role", "err", err)
		return nil, status.Error(codes.Internal, "failed to get updated role")
	}

	return &pb.UpdateRoleResponse{
		Success: true,
		Role:    role,
	}, nil
}

func (s *ServerGRPC) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	log := s.log.With(
		zap.String("name", req.Name),
		zap.String("company_id", req.CompanyId),
	)
	log.Debug("DeleteRole called")

	if err := s.logic.DeleteRole(ctx, req.Name); err != nil {
		log.Errorw("failed to delete role", "err", err)
		return nil, status.Error(codes.Internal, "failed to delete role")
	}

	return &pb.DeleteRoleResponse{
		Success: true,
	}, nil
}
