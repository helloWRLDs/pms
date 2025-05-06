package grpchandler

import (
	"context"

	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (s *ServerGRPC) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (res *pb.CreateTaskResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.CreateTaskResponse)
	res.Success = false

	created_id, err := s.logic.CreateTask(ctx, req.GetCreation())
	if err != nil {
		return res, err
	}

	created, err := s.logic.GetTask(ctx, created_id)
	if err != nil {
		return res, nil
	}

	res.CreatedTask = created
	res.Success = true
	return res, nil
}

func (s *ServerGRPC) GetTask(ctx context.Context, req *pb.GetTaskRequest) (res *pb.GetTaskResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.GetTaskResponse)
	res.Success = false

	task, err := s.logic.GetTask(ctx, req.GetId())
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Task = task

	return res, nil
}

func (s *ServerGRPC) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (res *pb.ListTasksResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.ListTasksResponse)
	res.Success = false

	filter := list.Filters{
		Pagination: list.Pagination{
			Page:    int(req.GetPage()),
			PerPage: int(req.GetPerPage()),
		},
		Fields: map[string]string{
			"t.project_id": utils.If(req.GetProjectId() != "", req.GetProjectId(), ""),
			"t.sprint_id":  utils.If(req.GetSprintId() != "", req.GetSprintId(), ""),
			"a.user_id":    utils.If(req.GetAssigneeId() != "", req.GetAssigneeId(), ""),
		},
	}

	tasks, err := s.logic.ListTasks(ctx, filter)
	if err != nil {
		return res, err
	}

	res.Tasks = &dto.TaskList{
		Page:       int32(tasks.Page),
		PerPage:    int32(tasks.PerPage),
		TotalPages: int32(tasks.TotalPages),
		TotalItems: int32(tasks.TotalItems),
		Items:      tasks.Items,
	}
	res.Success = true

	return res, nil
}
