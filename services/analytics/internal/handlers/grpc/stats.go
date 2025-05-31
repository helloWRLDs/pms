package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) GetUserTaskStats(ctx context.Context, req *pb.GetUserTaskStatsRequest) (res *pb.GetUserTaskStatsResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.GetUserTaskStatsResponse)
	res.Success = false

	stats, err := s.logic.GetUserTaskStats(ctx, req.CompanyId)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.Stats = stats

	return res, nil
}
