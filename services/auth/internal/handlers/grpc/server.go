package grpchandler

import (
	"context"

	"go.uber.org/zap"
	"pms.auth/internal/logic"
	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

type ServerGRPC struct {
	pb.UnimplementedAuthServiceServer

	logic *logic.Logic
	log   *zap.SugaredLogger
}

func New(logic *logic.Logic, log *zap.SugaredLogger) *ServerGRPC {
	return &ServerGRPC{
		logic: logic,
		log:   log,
	}
}

func (s *ServerGRPC) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (res *pb.LoginUserResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	payload, err := s.logic.LoginUser(ctx, req.Credentials)
	if err != nil {
		return nil, err
	}
	res = new(pb.LoginUserResponse)
	res.Success = true
	res.Payload = payload
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
