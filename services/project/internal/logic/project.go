package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
	projectdata "pms.project/internal/data/project"
)

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

	new := projectdata.Project{
		ID:          uuid.NewString(),
		Title:       creation.Title,
		Description: creation.Description,
		CompanyID:   creation.CompanyId,
		Status:      consts.PROJECT_STATUS_ACTIVE,
		CodeName:    creation.CodeName,
		CodePrefix:  utils.Ptr(creation.CodePrefix),
	}

	if err = l.Repo.Project.Create(tx, new); err != nil {
		log.Errorw("failed to create project", "err", err)
		return err
	}
	return nil
}

func (l *Logic) getProject(ctx context.Context, projectID string) (project *dto.Project, err error) {
	log := l.log.With(
		zap.String("func", "getTask"),
		zap.String("task_id", projectID),
	)
	log.Debug("getTask called")

	t, err := l.Repo.Project.GetByID(ctx, projectID)
	if err != nil {
		log.Errorw("failed to get task", "err", err)
		return nil, err
	}

	project = t.DTO()
	project.TotalTasks = int32(l.Repo.Task.Count(ctx, &dto.TaskFilter{
		ProjectId: projectID,
	}))
	project.DoneTasks = int32(l.Repo.Task.Count(ctx, &dto.TaskFilter{
		ProjectId: projectID,
		Status:    string(consts.TASK_STATUS_DONE),
	}))

	return project, nil
}

func (l *Logic) ListProjects(ctx context.Context, filter *dto.ProjectFilter) (result list.List[*dto.Project], err error) {
	log := l.log.With(
		zap.String("func", "ListProjects"),
		zap.String("filter", filter.String()),
	)
	log.Debug("ListProjects called")

	projects, err := l.Repo.Project.List(ctx, filter)
	if err != nil {
		log.Errorw("failed to get list of projects", "err", err)
		return list.List[*dto.Project]{}, err
	}
	log.Infow("fetched project list", "projects", projects)

	result = list.List[*dto.Project]{
		Items: make([]*dto.Project, len(projects.Items)),
		Pagination: list.Pagination{
			Page:       projects.Page,
			PerPage:    projects.PerPage,
			TotalItems: projects.TotalItems,
			TotalPages: projects.TotalPages,
		},
	}
	for i, p := range projects.Items {
		current, err := l.getProject(ctx, p.ID)
		if err != nil {
			log.Errorw("failed to get project", "err", err)
		}
		result.Items[i] = current
	}

	return result, nil
}

func (l *Logic) GetProjectByID(ctx context.Context, id string) (project *dto.Project, err error) {
	log := l.log.With(
		zap.String("func", "GetProjectByID"),
		zap.String("id", id),
	)
	log.Debug("GetProjectByID called")

	project, err = l.getProject(ctx, id)
	if err != nil {
		log.Errorw("failed to get project", "err", err)
		return nil, err
	}

	return project, nil
}
