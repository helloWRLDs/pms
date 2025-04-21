package grpchandler

import (
	"context"
	"strings"

	"pms.pkg/errs"
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
