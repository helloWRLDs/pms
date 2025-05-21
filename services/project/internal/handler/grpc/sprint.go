package grpchandler

import (
	"context"

	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) UpdateSprint(ctx context.Context, req *pb.UpdateSprintRequest) (res *pb.UpdateSprintResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.UpdateSprintResponse)
	res.Success = false

	updated, err := s.logic.UpdateSprint(ctx, req.Id, req.UpdatedSprint)
	if err != nil {
		return res, err
	}

	res.Success = true
	res.UpdatedSprint = updated

	return res, nil
}

func (s *ServerGRPC) GetSprint(ctx context.Context, req *pb.GetSprintRequest) (res *pb.GetSprintResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.GetSprintResponse)
	res.Success = false

	sprint, err := s.logic.GetSprint(ctx, req.Id)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.Sprint = sprint
	return res, nil
}

func (s *ServerGRPC) CreateSprint(ctx context.Context, req *pb.CreateSprintRequest) (res *pb.CreateSprintResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.CreateSprintResponse)
	res.Success = false

	created, err := s.logic.CreateSprint(ctx, req.Creation)
	if err != nil {
		return res, err
	}

	res.Success = true
	res.CreatedSprint = created

	return res, nil
}

func (s *ServerGRPC) ListSprints(ctx context.Context, req *pb.ListSprintsRequest) (res *pb.ListSprintsResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.ListSprintsResponse)
	res.Success = false

	sprintList, err := s.logic.ListSprints(ctx, req.Filter)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.Sprints = &dto.SprintList{
		Items:      sprintList.Items,
		Page:       int32(sprintList.Page),
		PerPage:    int32(sprintList.PerPage),
		TotalItems: int32(sprintList.TotalItems),
		TotalPages: int32(sprintList.TotalPages),
	}
	return res, nil
}
