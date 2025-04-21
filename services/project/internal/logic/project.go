package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	"pms.project/internal/entity"
)

func (l *Logic) CreateProject(ctx context.Context, creation *dto.ProjectCreation) error {
	log := l.log.With(
		zap.String("func", "CreateProject"),
		zap.Any("project_creation", creation),
	)
	log.Debug("CreateProject called")

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		log.Errorw("failed to start tx", "err", err)
		return err
	}
	defer func() {
		log.Debugw("err check", "isNil", err == nil, "err", err)
		l.Repo.EndTx(tx, err)
	}()

	new := entity.Project{
		ID:          uuid.New(),
		Title:       creation.Title,
		Description: creation.Description,
		CompanyID:   creation.CompanyId,
	}

	if err = l.Repo.Project.Create(tx, new); err != nil {
		log.Errorw("failed to create project", "err", err)
		return err
	}
	return nil
}
