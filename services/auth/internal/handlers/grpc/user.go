package grpchandler

import (
	"context"
	"strings"

	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) GetUser(ctx context.Context, req *pb.GetUserRequest) (res *pb.GetUserResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.GetUserResponse)
	res.Success = false

	if strings.Trim(req.UserID, " ") == "" {
		return res, errs.ErrBadGateway{
			Object: "user_id",
		}
	}

	profile, err := s.logic.GetProfile(ctx, req.UserID)
	if err != nil {
		return res, err
	}
	res.User = profile
	return res, nil
}

func (s *ServerGRPC) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (res *pb.ListUsersResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.ListUsersResponse)
	res.Success = false

	users, err := s.logic.ListUsers(ctx, req.Filter)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.Companies = &dto.UserList{
		Items:      users.Items,
		TotalPages: int32(users.TotalPages),
		Page:       int32(users.Page),
		PerPage:    int32(users.PerPage),
		TotalItems: int32(users.TotalItems),
	}
	return res, nil
}

func (s *ServerGRPC) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (res *pb.UpdateUserResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.UpdateUserResponse)
	res.Success = false

	if strings.Trim(req.Id, " ") == "" {
		return res, errs.ErrBadGateway{
			Object: "user_id",
		}
	}

	updated, err := s.logic.UpdateUser(ctx, req.Id, req.UpdatedUser)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.User = updated
	return res, nil
}
