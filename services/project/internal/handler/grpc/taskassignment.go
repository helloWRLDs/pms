package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) TaskAssign(ctx context.Context, req *pb.TaskAssignRequest) (res *pb.TaskAssignResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.TaskAssignResponse)
	res.Success = false

	if err := s.logic.AssignTask(ctx, req.UserId, req.TaskId); err != nil {
		return res, err
	}
	res.Success = true
	return res, nil
}

func (s *ServerGRPC) TaskUnassign(ctx context.Context, req *pb.TaskUnassignRequest) (res *pb.TaskUnassignResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.TaskUnassignResponse)
	res.Success = false

	if err := s.logic.UnassignTask(ctx, req.UserId, req.TaskId); err != nil {
		return res, err
	}
	res.Success = true
	return res, nil
}
