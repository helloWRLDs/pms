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

	payload, err := s.logic.LoginUser(ctx, req.Credentials)
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
	if err := s.logic.RegisterUser(ctx, req.NewUser); err != nil {
		log.Errorw("failed to register user", "err", err)
		return res, err
	}
	created, err := s.logic.Repo.User.GetByEmail(ctx, req.NewUser.Email)
	if err != nil {
		return nil, err
	}
	res.Success = true
	res.User = created.DTO()
	return res, nil
}
