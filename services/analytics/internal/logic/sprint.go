package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) getSprint(ctx context.Context, sprintID string) (*dto.Sprint, error) {
	log := l.log.Named("getSprint").With(
		zap.Any("sprint_id", sprintID),
	)
	log.Debug("getSprint called")

	sprintRes, err := l.projectClient.GetSprint(ctx, &pb.GetSprintRequest{
		Id: sprintID,
	})
	if err != nil {
		log.Error("failed to get sprint", zap.Error(err))
		return nil, err
	}
	tasksRes, err := l.projectClient.ListTasks(ctx, &pb.ListTasksRequest{
		Filter: &dto.TaskFilter{
			SprintId: sprintID,
			Page:     1,
			PerPage:  10000,
		},
	})
	if err != nil {
		log.Error("failed to list tasks", zap.Error(err))
		return nil, err
	}
	sprint := sprintRes.Sprint
	sprint.Tasks = tasksRes.Tasks.Items
	return sprint, nil
}
