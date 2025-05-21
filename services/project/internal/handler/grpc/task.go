package grpchandler

import (
	"context"

	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
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

	tasks, err := s.logic.ListTasks(ctx, req.Filter)
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

func (s *ServerGRPC) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (res *pb.UpdateTaskResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.UpdateTaskResponse)
	res.Success = false

	if err = s.logic.UpdateTask(ctx, req.Id, req.UpdatedTask); err != nil {
		return res, err
	}

	res.Success = true
	return res, nil
}
