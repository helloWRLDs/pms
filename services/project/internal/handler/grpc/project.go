package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (res *pb.CreateProjectResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.CreateProjectResponse)
	res.Success = false

	if err = s.logic.CreateProject(ctx, req.Creation); err != nil {
		return res, err
	}

	return res, nil
}
