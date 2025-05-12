package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (l *Logic) GetTask(ctx context.Context, taskID string) (*dto.Task, error) {
	log := l.log.With(
		zap.String("func", "GetTask"),
	)
	log.Debug("GetTask called")

	session, err := l.GetSessionInfo(ctx)
	if err != nil {
		log.Errorw("failed to get session", "err", err)
		return nil, err
	}
	log.Debug("session retrieved", "session", session)

	res, err := l.projectClient.GetTask(ctx, &pb.GetTaskRequest{Id: taskID})
	if err != nil {
		log.Errorw("failed to get task", "err", err)
		return nil, err
	}

	return res.Task, nil
}

func (l *Logic) CreateTask(ctx context.Context, creation *dto.TaskCreation) (err error) {
	log := l.log.With(
		zap.String("func", "CreateTask"),
		zap.Any("task_creation", creation),
	)
	log.Debug("CreateTask called")

	session, err := l.GetSessionInfo(ctx)
	if err != nil {
		log.Errorw("failed to get session", "err", err)
		return err
	}
	log.Debug("session retrieved", "session", session)
	res, err := l.projectClient.CreateTask(ctx, &pb.CreateTaskRequest{Creation: creation})
	if err != nil {
		log.Errorw("failed to create task", "err", err)
		return err
	}
	log.Debug("task created", "res", res)

	return nil
}

func (l *Logic) ListTasks(ctx context.Context, filter list.Filters) (result list.List[*dto.Task], err error) {
	log := l.log.With(
		zap.String("func", "ListTasks"),
	)
	log.Debug("ListTasks called")

	var (
		isExist = func(key string) bool {
			_, ok := filter.Fields[key]
			return ok
		}
		projectID  = utils.If(isExist("project_id"), filter.Fields["project_id"], "")
		sprintID   = utils.If(isExist("sprint_id"), filter.Fields["sprint_id"], "")
		assigneeID = utils.If(isExist("assignee_id"), filter.Fields["assignee_id"], "")
	)

	res, err := l.projectClient.ListTasks(ctx, &pb.ListTasksRequest{
		ProjectId:  projectID,
		SprintId:   sprintID,
		AssigneeId: assigneeID,
		Page:       int32(filter.Page),
		PerPage:    int32(filter.PerPage),
	})
	if err != nil {
		log.Errorw("failed to list tasks", "err", err)
		return list.List[*dto.Task]{}, err
	}

	return list.List[*dto.Task]{
		Items: res.Tasks.Items,
		Pagination: list.Pagination{
			Page:       int(res.Tasks.Page),
			PerPage:    int(res.Tasks.PerPage),
			TotalPages: int(res.Tasks.TotalPages),
			TotalItems: int(res.Tasks.TotalItems),
		},
	}, nil
}
