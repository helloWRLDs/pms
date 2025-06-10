package grpchandler

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (res *pb.CreateRoleResponse, err error) {
	log := s.log.With(
		zap.Any("request", req),
	)
	log.Debug("CreateRole called")

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.CreateRoleResponse)
	res.Success = false
	if err := s.logic.CreateRole(ctx, req.Role); err != nil {
		log.Errorw("failed to create role", "err", err)
		return nil, err
	}

	role, err := s.logic.GetRole(ctx, req.Role.Name)
	if err != nil {
		log.Errorw("failed to get created role", "err", err)
		return nil, err
	}

	res.Success = true
	res.Role = role
	return res, nil
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

func (s *ServerGRPC) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (res *pb.ListRolesResponse, err error) {
	log := s.log.With(
		zap.String("company_id", req.Filter.CompanyId),
	)
	log.Debug("ListRoles called")

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.ListRolesResponse)
	res.Success = false

	roles, err := s.logic.ListRoles(ctx, req.Filter)
	if err != nil {
		log.Errorw("failed to list roles", "err", err)
		return nil, err
	}
	res.Success = true
	res.Roles = roles
	return res, nil
}

func (s *ServerGRPC) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (res *pb.UpdateRoleResponse, err error) {
	log := s.log.With(
		zap.String("name", req.Name),
		zap.Any("updated_role", req.Role),
	)
	log.Debug("UpdateRole called")

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.UpdateRoleResponse)
	res.Success = false

	if err := s.logic.UpdateRole(ctx, req.Name, req.Role, req.CompanyId); err != nil {
		log.Errorw("failed to update role", "err", err)
		return nil, err
	}

	res.Success = true
	return res, nil
}

func (s *ServerGRPC) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (res *pb.DeleteRoleResponse, err error) {
	log := s.log.With(
		zap.String("name", req.Name),
		zap.String("company_id", req.CompanyId),
	)
	log.Debug("DeleteRole called")

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.DeleteRoleResponse)
	res.Success = false

	if err := s.logic.DeleteRole(ctx, req.Name, req.CompanyId); err != nil {
		log.Errorw("failed to delete role", "err", err)
		return nil, err
	}

	res.Success = true
	return res, nil
}
