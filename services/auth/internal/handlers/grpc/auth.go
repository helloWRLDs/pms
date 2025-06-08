package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (res *pb.LoginUserResponse, err error) {
	log := s.log.With("func", "LoginUser", "pkg", "grpchandler")
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	payload, err := s.logic.LoginUser(ctx, nil, req.Credentials)
	if err != nil {
		log.Errorw("failed to login user", "err", err)
		return nil, err
	}
	res = new(pb.LoginUserResponse)
	res.Success = true
	res.Payload = payload

	user, err := s.logic.GetProfile(ctx, payload.User.Id)
	if err != nil {
		log.Errorw("failed to get user profile", "err", err)
		return nil, err
	}
	res.Payload.User = user
	return res, nil
}

func (s *ServerGRPC) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (res *pb.RegisterUserResponse, err error) {
	log := s.log.With("func", "RegisterUser", "pkg", "grpchandler")
	defer func() {
		err = errs.WrapGRPC(err)
	}()
	res = new(pb.RegisterUserResponse)
	created, err := s.logic.RegisterUser(ctx, req.NewUser)
	if err != nil {
		log.Errorw("failed to register user", "err", err)
		return res, err
	}

	res.Success = true
	res.User = created
	return res, nil
}

func (s *ServerGRPC) GetUserRole(ctx context.Context, req *pb.GetUserRoleRequest) (res *pb.GetUserRoleResponse, err error) {
	log := s.log.Named("GetUserRole")
	log.Debug("GetUserRole called")

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = &pb.GetUserRoleResponse{
		Success: false,
	}

	// Get participant to find role name
	participant, err := s.logic.Repo.Participant.GetByUserAndCompany(ctx, req.UserId, req.CompanyId)
	if err != nil {
		log.Errorw("failed to get participant", "err", err)
		return nil, err
	}

	// Get role by name
	role, err := s.logic.Repo.Role.GetByName(ctx, participant.Role)
	if err != nil {
		log.Errorw("failed to get role", "err", err)
		return nil, err
	}

	res.Success = true
	res.Role = role.DTO()

	return res, nil
}
