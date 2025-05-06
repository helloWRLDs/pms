package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (l *Logic) ListProjects(ctx context.Context, company_id string, filter list.Filters) (result list.List[*dto.Project], err error) {
	log := l.log.With(
		zap.String("func", "ListProjects"),
		zap.String("filter", filter.String()),
	)
	log.Debug("ListProjects called")

	// currentSession, err := l.GetSessionInfo(ctx)
	// if err != nil {
	// 	return list.List[*dto.Project]{}, err
	// }

	res, err := l.projectClient.ListProjects(ctx, &pb.ListProjectsRequest{
		Page:      int32(filter.Page),
		PerPage:   int32(filter.PerPage),
		CompanyId: company_id,
	})
	if err != nil {
		log.Errorw("failed to list projects", "err", err)
		return list.List[*dto.Project]{}, err
	}

	return list.List[*dto.Project]{
		Items: res.Projects.Items,
		Pagination: list.Pagination{
			Page:       int(res.Projects.Page),
			PerPage:    int(res.Projects.PerPage),
			TotalPages: int(res.Projects.TotalPages),
			TotalItems: int(res.Projects.TotalItems),
		},
	}, nil
}

func (l *Logic) GetProjectByID(ctx context.Context, projectID string) (*dto.Project, error) {
	log := l.log.With(
		zap.String("func", "GetProjectByID"),
	)
	log.Debug("GetProjectByID called")

	currentSession, err := l.GetSessionInfo(ctx)
	if err != nil {
		return nil, err
	}
	print(utils.JSON(currentSession))

	res, err := l.projectClient.GetProject(ctx, &pb.GetProjectRequest{Id: projectID})
	if err != nil {
		log.Errorw("failed to get project", "err", err)
		return nil, err
	}

	return res.Project, nil
}

func (l *Logic) CreateProject(ctx context.Context, creation *dto.ProjectCreation) (err error) {
	log := l.log.With(
		zap.String("func", "CreateProject"),
		zap.Any("project_creation", creation),
	)
	log.Debug("CreateProject called")

	session, err := l.GetSessionInfo(ctx)
	if err != nil {
		log.Errorw("failed to get session", "err", err)
		return err
	}
	log.Debug("session retrieved", "session", session)

	res, err := l.projectClient.CreateProject(ctx, &pb.CreateProjectRequest{Creation: creation})
	if err != nil {
		log.Errorw("failed to create project", "err", err)
		return err
	}
	log.Debug("project created", "res", res)

	return nil
}
