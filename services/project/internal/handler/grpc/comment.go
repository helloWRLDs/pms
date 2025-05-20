package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) CreateTaskComment(ctx context.Context, req *pb.CreateTaskCommentRequest) (res *pb.CreateTaskCommentResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.CreateTaskCommentResponse)
	res.Success = false

	created, err := s.logic.CreateTaskComment(ctx, req.Creation)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.CreatedComment = created

	return res, nil
}

func (s *ServerGRPC) ListTaskComments(ctx context.Context, req *pb.ListTaskCommentsRequest) (res *pb.ListTaskCommentsResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.ListTaskCommentsResponse)
	res.Success = false

	list, err := s.logic.ListTaskComments(ctx, req.Filter)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.List = list

	return res, nil
}

func (s *ServerGRPC) GetTaskComment(ctx context.Context, req *pb.GetTaskCommentRequest) (res *pb.GetTaskCommentResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.GetTaskCommentResponse)
	res.Success = false

	comment, err := s.logic.GetTaskComment(ctx, req.Id)
	if err != nil {
		return res, nil
	}
	res.Success = true
	res.Comment = comment

	return res, nil
}

func (s *ServerGRPC) DeleteTaskComment(context.Context, *pb.DeleteTaskCommentRequest) (res *pb.DeleteTaskCommentResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.DeleteTaskCommentResponse)
	res.Success = false

	return res, nil
}
