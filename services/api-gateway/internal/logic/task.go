package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	notifiermq "pms.pkg/transport/mq/notifier"
	"pms.pkg/type/list"
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

func (l *Logic) DeleteTask(ctx context.Context, taskID string) (err error) {
	log := l.log.With(
		zap.String("func", "DeleteTask"),
		zap.String("task_id", taskID),
	)
	log.Debug("DeleteTask called")

	res, err := l.projectClient.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: taskID})
	if err != nil {
		log.Errorw("failed to delete task", "err", err)
		return err
	}
	log.Debug("task deleted", "res", res)

	return nil
}

func (l *Logic) UpdateTask(ctx context.Context, taskID string, task *dto.Task) (err error) {
	log := l.log.With(
		zap.String("func", "UpdateTask"),
		zap.String("task_id", taskID),
		zap.Any("updated_task", task),
	)
	log.Debug("UpdateTask called")

	// session, err := l.GetSessionInfo(ctx)
	// if err != nil {
	// 	return err
	// }

	// if !utils.ContainsInArray(session.Companies, task.ProjectId) {
	// 	return errs.ErrNotFound{
	// 		Object: "task",
	// 		Value:  taskID,
	// 	}
	// }

	res, err := l.projectClient.UpdateTask(ctx, &pb.UpdateTaskRequest{
		Id:          taskID,
		UpdatedTask: task,
	})

	if err != nil {
		log.Errorw("failed to update task", "err", err)
		return err
	}
	log.Debug("task updated", "res", res)

	return nil
}

func (l *Logic) ListTasks(ctx context.Context, filter *dto.TaskFilter) (result list.List[*dto.Task], err error) {
	log := l.log.With(
		zap.String("func", "ListTasks"),
	)
	log.Debug("ListTasks called")

	res, err := l.projectClient.ListTasks(ctx, &pb.ListTasksRequest{
		Filter: filter,
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

func (l *Logic) TaskAssign(ctx context.Context, taskID, userID string) error {
	log := l.log.Named("TaskAssign").With(
		zap.String("task_id", taskID),
		zap.String("user_id", userID),
	)
	log.Info("TaskAssign called")
	assignRes, err := l.projectClient.TaskAssign(ctx, &pb.TaskAssignRequest{
		TaskId: taskID,
		UserId: userID,
	})
	if err != nil {
		return err
	}
	log.Info("task assignment result", "res", assignRes)

	taskRes, err := l.projectClient.GetTask(ctx, &pb.GetTaskRequest{Id: taskID})
	if err != nil {
		log.Errorw("failed to get task details", "err", err)
		return err
	}

	userRes, err := l.authClient.GetUser(ctx, &pb.GetUserRequest{UserID: userID})
	if err != nil {
		log.Errorw("failed to get user details", "err", err)
		return err
	}

	projectRes, err := l.projectClient.GetProject(ctx, &pb.GetProjectRequest{Id: taskRes.Task.ProjectId})
	if err != nil {
		log.Errorw("failed to get project details", "err", err)
		return err
	}

	err = l.notificationMQ.Publish(ctx, notifiermq.TaskAssignmentMessage{
		MetaData: notifiermq.MetaData{
			ToEmail: userRes.User.Email,
		},
		AssigneeName: userRes.User.FirstName + " " + userRes.User.LastName,
		TaskName:     taskRes.Task.Title,
		TaskId:       taskID,
		ProjectName:  projectRes.Project.Title,
	})
	if err != nil {
		log.Errorw("failed to send task assignment notification", "err", err)
		return nil
	}

	return nil
}

func (l *Logic) TaskUnassign(ctx context.Context, taskID, userID string) error {
	log := l.log.Named("TaskUnassign").With(
		zap.String("task_id", taskID),
		zap.String("user_id", userID),
	)
	log.Info("TaskUnassign called")

	unassignRes, err := l.projectClient.TaskUnassign(ctx, &pb.TaskUnassignRequest{
		TaskId: taskID,
		UserId: userID,
	})
	log.Infow("task unassignment result", "res", unassignRes)
	if err != nil {
		return err
	}

	return nil

}
