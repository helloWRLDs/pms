package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
	sprintdata "pms.project/internal/data/sprint"
)

func (l *Logic) CreateSprint(ctx context.Context, creation *dto.SprintCreation) (created *dto.Sprint, err error) {
	log := l.log.Named("CreateSprint").With(
		zap.Any("sprint_creation", creation),
	)
	log.Debug("func", "CreateSprint")

	if err := l.createSprint(ctx, creation); err != nil {
		return nil, err
	}

	sprints, err := l.ListSprints(ctx, &dto.SprintFilter{
		ProjectId: creation.ProjectId,
		Title:     creation.Title,
	})
	if err != nil {
		return nil, err
	}
	if len(sprints.Items) > 0 {
		created = sprints.Items[0]
	}

	return created, nil
}

func (l *Logic) createSprint(ctx context.Context, creation *dto.SprintCreation) (err error) {
	log := l.log.Named("createSprint").With(
		zap.Any("sprint_creation", creation),
	)
	log.Debug("createSprint called")

	existing, err := l.Repo.Sprint.List(ctx, &dto.SprintFilter{
		ProjectId: creation.ProjectId,
		Title:     creation.Title,
	})
	if err == nil && len(existing.Items) > 0 {
		log.Error("Sprint already exists")
		return errs.ErrAlreadyExist{
			Object: "sprint",
			Field:  "title",
			Value:  creation.Title,
		}
	}
	newSprint := sprintdata.Sprint{
		Title:       creation.Title,
		Description: creation.Description,
		ProjectID:   creation.ProjectId,
		StartDate:   creation.StartDate.AsTime(),
		EndDate:     creation.EndDate.AsTime(),
	}
	if err := l.Repo.Sprint.Create(ctx, newSprint); err != nil {
		log.Error("CreateSprint failed", zap.Error(err))
		return err
	}

	return nil
}

func (l *Logic) ListSprints(ctx context.Context, filter *dto.SprintFilter) (res list.List[*dto.Sprint], err error) {
	log := l.log.Named("ListSprints").With(
		zap.String("filter", filter.String()),
	)
	log.Debug("ListSprints called")

	sprints, err := l.Repo.Sprint.List(ctx, filter)
	if err != nil {
		log.Error("ListSprints failed", zap.Error(err))
		return list.List[*dto.Sprint]{}, err
	}

	for _, s := range sprints.Items {
		dtoSprint := s.DTO()
		res.Items = append(res.Items, dtoSprint)
	}
	res.TotalItems = sprints.TotalItems
	res.TotalPages = sprints.TotalPages
	res.Page = sprints.Page
	res.PerPage = sprints.PerPage

	return res, nil
}

func (l *Logic) GetSprint(ctx context.Context, sprintID string) (sprint *dto.Sprint, err error) {
	log := l.log.Named("GetSprint").With(
		zap.String("sprint_id", sprintID),
	)
	log.Debug("GetSprint called")

	s, err := l.Repo.Sprint.GetByID(ctx, sprintID)
	if err != nil {
		log.Error("GetSprintByID failed", zap.Error(err))
		return nil, err
	}

	return s.DTO(), nil
}

func (l *Logic) UpdateSprint(ctx context.Context, sprintID string, sprint *dto.Sprint) (updated *dto.Sprint, err error) {
	log := l.log.Named("UpdateSprint").With(
		zap.Any("sprint", sprint),
	)
	log.Debug("ListSprints called")

	if err := l.Repo.Sprint.Update(ctx, sprintID, utils.Value(sprintdata.Entity(sprint))); err != nil {
		log.Error("UpdateSprint failed", zap.Error(err))
		return nil, err
	}

	updated, err = l.GetSprint(ctx, sprintID)
	if err != nil {
		log.Error("GetSprintByID failed", zap.Error(err))
		return nil, err
	}

	return updated, nil
}
