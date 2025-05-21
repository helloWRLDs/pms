package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) CreateSprint(ctx context.Context, sprint *dto.SprintCreation) (*dto.Sprint, error) {
	log := l.log.Named("CreateSprint").With(
		zap.Any("sprint", sprint),
	)
	log.Debug("CreateSprint called")

	sprintRes, err := l.projectClient.CreateSprint(ctx, &pb.CreateSprintRequest{
		Creation: sprint,
	})
	if err != nil {
		log.Errorw("failed to create sprint", "err", err)
		return nil, err
	}
	return sprintRes.CreatedSprint, nil
}

func (l *Logic) ListSprints(ctx context.Context, filter *dto.SprintFilter) (*dto.SprintList, error) {
	log := l.log.Named("ListSprints").With(
		zap.Any("filter", filter),
	)
	log.Debug("ListTasks called")

	sprintRes, err := l.projectClient.ListSprints(ctx, &pb.ListSprintsRequest{
		Filter: filter,
	})
	if err != nil {
		log.Errorw("failed to list sprints", "err", err)
		return nil, err
	}

	return sprintRes.Sprints, nil
}

func (l *Logic) GetSprint(ctx context.Context, sprintID string) (*dto.Sprint, error) {
	log := l.log.Named("GetSprint").With(
		zap.String("sprint_id", sprintID),
	)
	log.Debug("GetSprint called")

	sprintRes, err := l.projectClient.GetSprint(ctx, &pb.GetSprintRequest{
		Id: sprintID,
	})
	if err != nil {
		log.Errorw("failed to get sprint", "err", err)
		return nil, err
	}

	return sprintRes.Sprint, nil
}

func (l *Logic) UpdateSprint(ctx context.Context, sprintID string, updated *dto.Sprint) (*dto.Sprint, error) {
	log := l.log.Named("UpdateSprint").With(
		zap.String("sprint_id", sprintID),
	)
	log.Debug("UpdateSprint called")

	updateRes, err := l.projectClient.UpdateSprint(ctx, &pb.UpdateSprintRequest{
		Id:            sprintID,
		UpdatedSprint: updated,
	})
	if err != nil {
		log.Errorw("failed to update sprint", "err", err)
		return nil, err
	}

	return updateRes.UpdatedSprint, nil
}
