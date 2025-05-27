package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) ListCompanies(ctx context.Context, filter *dto.CompanyFilter) (*dto.CompanyList, error) {
	log := l.log.With(
		zap.String("func", "ListCompanies"),
	)
	log.Info("ListCompanies called")

	session, err := l.GetSessionInfo(ctx)
	if err != nil {
		log.Errorw("failed to get session", "err", err)
		return nil, err
	}
	log.Infow("retrieved session", "session", session)
	companyRes, err := l.authClient.ListCompanies(ctx, &pb.ListCompaniesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	for i, comp := range companyRes.Companies.Items {
		projectRes, err := l.projectClient.ListProjects(ctx, &pb.ListProjectsRequest{
			Filter: &dto.ProjectFilter{
				Page:      1,
				PerPage:   100,
				CompanyId: comp.Id,
			},
		})
		if err != nil {
			log.Errorw("failed to list projects", "err", err)
		}
		log.Infow("retrieved projects", "projects", projectRes.Projects)
		companyRes.Companies.Items[i].Projects = projectRes.Projects
	}

	return companyRes.Companies, nil
}

func (l *Logic) CompanyRemoveParticipant(ctx context.Context, companyID, userID string) error {
	log := l.log.Named("CompanyRemoveParticipant").With(
		zap.String("company_id", companyID),
		zap.String("user_id", userID),
	)
	log.Debug("CompanyRemoveParticipant called")

	res, err := l.authClient.RemoveParticipant(ctx, &pb.RemoveParticipantRequest{
		CompanyId: companyID,
		UserId:    userID,
	})
	log.Infow("trying to remove participant", "res", res)

	if err != nil {
		log.Errorw("failed to remove participant", "err", err)
		return err
	}

	return nil
}

func (l *Logic) CompanyAddParticipant(ctx context.Context, companyID, userID string) error {
	log := l.log.Named("CompanyAddParticipant").With(
		zap.Any("user_id", userID),
		zap.String("company_id", companyID),
	)
	log.Debug("CompanyAddParticipant called")

	res, err := l.authClient.AddParticipant(ctx, &pb.AddParticipantRequest{
		CompanyId: companyID,
		UserId:    userID,
		RoleId:    "admin",
	})
	log.Infow("adding participant", "res", res)

	if err != nil {
		log.Errorw("failed to add participant", "err", err)
		return err
	}

	return nil
}

func (l *Logic) CreateCompany(ctx context.Context, creation *dto.NewCompany) error {
	log := l.log.Named("CreateCompany").With(
		zap.Any("creation", creation),
	)
	log.Debug("GetCompany called")

	sessionInfo, err := l.GetSessionInfo(ctx)
	if err != nil {
		return err
	}
	creationRes, err := l.authClient.CreateCompany(ctx, &pb.CreateCompanyRequest{
		Company: creation,
		UserId:  sessionInfo.UserID,
	})
	if err != nil {
		return err
	}
	log.Infow("created company", "created", creationRes.Company)

	return nil
}

func (l *Logic) GetCompany(ctx context.Context, companyID string) (*dto.Company, error) {
	log := l.log.With(
		zap.String("func", "GetCompany"),
		zap.String("company_id", companyID),
	)
	log.Debug("GetCompany called")

	company := new(dto.Company)

	_, err := l.GetSessionInfo(ctx)
	if err != nil {
		return nil, err
	}

	companyRes, err := l.authClient.GetCompany(ctx, &pb.GetCompanyRequest{
		Id: companyID,
	})
	if err != nil {
		return nil, err
	}
	company = companyRes.Company

	projectsRes, err := l.projectClient.ListProjects(ctx, &pb.ListProjectsRequest{
		Filter: &dto.ProjectFilter{
			Page:      1,
			PerPage:   10,
			CompanyId: companyID,
		},
	})
	if err != nil {
		return nil, err
	}
	company.Projects = projectsRes.Projects

	return company, nil
}
