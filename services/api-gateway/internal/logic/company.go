package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/type/list"
)

func (l *Logic) ListCompanies(ctx context.Context, filter list.Filters) (*dto.CompanyList, error) {
	log := l.log.With(
		zap.String("func", "ListCompanies"),
	)
	log.Debug("ListCompanies called")

	session, err := l.GetSessionInfo(ctx)
	if err != nil {
		log.Errorw("failed to get session", "err", err)
		return nil, err
	}
	log.Infow("retrieved session", "session", session)
	res, err := l.authClient.ListCompanies(ctx, &pb.ListCompaniesRequest{
		UserId:  session.UserID,
		Page:    int32(filter.Page),
		PerPage: int32(filter.PerPage),
	})
	if err != nil {
		return nil, err
	}

	return res.Companies, nil
}

func (l *Logic) GetCompany(ctx context.Context, companyID string) (*dto.Company, error) {
	log := l.log.With(
		zap.String("func", "GetCompany"),
		zap.String("company_id", companyID),
	)
	log.Debug("GetCompany called")

	_, err := l.GetSessionInfo(ctx)
	if err != nil {
		return nil, err
	}

	res, err := l.authClient.GetCompany(ctx, &pb.GetCompanyRequest{
		Id: companyID,
	})
	if err != nil {
		return nil, err
	}
	return res.Company, nil
}
