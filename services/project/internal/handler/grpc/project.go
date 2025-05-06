package grpchandler

import (
	"context"

	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/type/list"
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

func (s *ServerGRPC) GetProject(ctx context.Context, req *pb.GetProjectRequest) (res *pb.GetProjectResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.GetProjectResponse)
	res.Success = false

	project, err := s.logic.GetProjectByID(ctx, req.GetId())
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Project = project

	return res, nil
}

func (s *ServerGRPC) ListProjects(ctx context.Context, req *pb.ListProjectsRequest) (res *pb.ListProjectsResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.ListProjectsResponse)
	res.Success = false

	list, err := s.logic.ListProjects(ctx, list.Filters{
		Pagination: list.Pagination{
			Page:    int(req.Page),
			PerPage: int(req.PerPage),
		},
		Fields: map[string]string{
			"p.company_id": req.CompanyId,
		},
	})
	res.Success = true
	res.Projects = &dto.ProjectList{
		Items:      list.Items,
		Page:       int32(list.Page),
		PerPage:    int32(list.PerPage),
		TotalItems: int32(list.TotalItems),
		TotalPages: int32(list.TotalPages),
	}

	return res, nil
}
