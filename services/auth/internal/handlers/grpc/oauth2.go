package grpchandler

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) CompleteOAuth2(ctx context.Context, req *pb.CompleteOAuth2Request) (res *pb.CompleteOAuth2Response, err error) {
	log := s.log.Named("CompleteOAuth2")
	log.Debug("CompleteOAuth2 called")

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = &pb.CompleteOAuth2Response{
		Success: false,
	}

	user, payload, err := s.logic.CompleteOAuth2(ctx, req.Provider, req.Code)
	if err != nil {
		log.Error("failed to complete OAuth2", zap.Error(err))
		return nil, err
	}

	res.Success = true
	res.User = user
	res.Payload = payload

	return res, nil
}

func (s *ServerGRPC) InitiateOAuth2(ctx context.Context, req *pb.InitiateOAuth2Request) (res *pb.InitiateOAuth2Response, err error) {
	log := s.log.Named("InitiateOAuth2")
	log.Debug("InitiateOAuth2 called")

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	authURL, err := s.logic.InitiateOAuth2(req.Provider)
	if err != nil {
		log.Error("failed to initiate OAuth2", zap.Error(err))
		return nil, err
	}

	return &pb.InitiateOAuth2Response{
		AuthUrl: authURL,
	}, nil
}
