package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
	"pms.project/internal/data/models"
)

func (l *Logic) CalculateProjectCompletion(ctx context.Context, projectID string) (err error) {
	log := l.log.With(
		zap.String("func", "CalculateProjectCompletion"),
		zap.String("project_id", projectID),
	)
	log.Debug("CalculateProjectCompletion called")

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		log.Errorw("failed to start tx", "err", err)
		return err
	}
	defer func() {
		log.Debugw("err check", "isNil", err == nil, "err", err)
		l.Repo.EndTx(tx, err)
	}()

	project, err := l.Repo.Project.GetByID(tx, projectID)
	if err != nil {
		log.Errorw("failed to get project", "err", err)
		return err
	}
	log.Debugw("retrieved project", "project", project)

	tasks, err := l.Repo.Task.List(tx, list.Filters{
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10000,
		},
		Fields: map[string]string{
			"p.project_id": projectID,
		},
	})
	if err != nil {
		log.Errorw("failed to list tasks", "err", err)
		return err
	}

	if len(tasks.Items) == 0 {
		log.Info("no tasks found")
		project.Progress = utils.Ptr(0)
	} else {
		completedCount := 0
		for _, task := range tasks.Items {
			if task.Status == consts.TASK_STATUS_DONE {
				completedCount++
			}
		}
		project.Progress = utils.Ptr((completedCount * 100) / len(tasks.Items))
	}

	if err = l.Repo.Project.Update(tx, project.ID.String(), project); err != nil {
		log.Errorw("failed to update project", "err", err)
		return err
	}
	log.Infow("recalculated progression", "progress", project.Progress)
	return nil
}

func (l *Logic) CreateProject(ctx context.Context, creation *dto.ProjectCreation) (err error) {
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

	new := models.Project{
		ID:          uuid.New(),
		Title:       creation.Title,
		Description: creation.Description,
		CompanyID:   creation.CompanyId,
		Status:      consts.PROJECT_STATUS_ACTIVE,
	}

	if err = l.Repo.Project.Create(tx, new); err != nil {
		log.Errorw("failed to create project", "err", err)
		return err
	}
	return nil
}

func (l *Logic) ListProjects(ctx context.Context, filter list.Filters) (result list.List[*dto.Project], err error) {
	log := l.log.With(
		zap.String("func", "ListProjects"),
		zap.String("filter", filter.String()),
	)
	log.Debug("ListProjects called")

	projects, err := l.Repo.Project.List(ctx, filter)
	if err != nil {
		return list.List[*dto.Project]{}, err
	}

	result = list.List[*dto.Project]{
		Items: func() []*dto.Project {
			res := make([]*dto.Project, len(projects.Items))
			for i, p := range projects.Items {
				res[i] = &dto.Project{
					Id:          p.ID.String(),
					Title:       p.Title,
					Description: p.Description,
					CompanyId:   p.CompanyID,
					Status:      string(p.Status),
					CreatedAt:   timestamppb.New(p.CreatedAt.Time),
					UpdatedAt:   timestamppb.New(p.UpdatedAt.Time),
					TotalTasks: func() int32 {
						return int32(l.Repo.Task.Count(ctx, list.Filters{
							Fields: map[string]string{
								"p.id": p.ID.String(),
							},
						}))
					}(),
				}
			}
			return res
		}(),
		Pagination: list.Pagination{
			Page:       projects.Page,
			PerPage:    projects.PerPage,
			TotalItems: projects.TotalItems,
			TotalPages: projects.TotalPages,
		},
	}

	return result, nil
}

func (l *Logic) GetProjectByID(ctx context.Context, id string) (project *dto.Project, err error) {
	log := l.log.With(
		zap.String("func", "GetProjectByID"),
		zap.String("id", id),
	)
	log.Debug("GetProjectByID called")

	if err := l.CalculateProjectCompletion(ctx, id); err != nil {
		log.Debugw("failed to calculate progress for project", "err", err)
	}

	p, err := l.Repo.Project.GetByID(ctx, id)
	if err != nil {
		log.Errorw("failed to get project by id", "err", err)
		return nil, err
	}

	project = &dto.Project{
		Id:                 p.ID.String(),
		Title:              p.Title,
		Description:        p.Description,
		CompanyId:          p.CompanyID,
		CreatedAt:          timestamppb.New(p.CreatedAt.Time),
		UpdatedAt:          timestamppb.New(p.UpdatedAt.Time),
		Status:             string(p.Status),
		CompletionProgress: int32(utils.Value(p.Progress)),
		TotalTasks: func() int32 {
			return int32(l.Repo.Task.Count(ctx, list.Filters{
				Fields: map[string]string{
					"project_id": p.ID.String(),
				},
			}))
		}(),
	}

	return project, nil
}
