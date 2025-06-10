package grpchandler

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyRequest) (res *pb.UpdateCompanyResponse, err error) {
	log := s.log.With("func", "RemoveParticipant", "pkg", "grpchandler")
	log.Debug("RemoveParticipant called")

	res = new(pb.UpdateCompanyResponse)
	res.Success = false

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	if err := s.logic.UpdateCompany(ctx, req.GetId(), req.GetCompany()); err != nil {
		log.Errorw("failed to remove participant", "err", err)
		return nil, err
	}
	updated, err := s.logic.GetCompany(ctx, req.GetId())
	if err != nil {
		log.Errorw("failed to get updated company", "err", err)
		return nil, err
	}
	res.Success = true
	res.Company = updated
	return
}

func (s *ServerGRPC) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (res *pb.CreateCompanyResponse, err error) {
	log := s.log.With("func", "ListCompanies", "pkg", "grpchandler")
	log.Debug("ListCompanies called")

	res = new(pb.CreateCompanyResponse)
	res.Success = false

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	created, err := s.logic.CreateCompany(ctx, req.UserId, req.Company)
	if err != nil {
		log.Errorw("failed to create company", "err", err)
		return nil, err
	}
	res.Success = true
	res.Company = created
	return
}

func (s *ServerGRPC) ListCompanies(ctx context.Context, req *pb.ListCompaniesRequest) (res *pb.ListCompaniesResponse, err error) {
	log := s.log.With("func", "ListCompanies", "pkg", "grpchandler")
	log.Debug("ListCompanies called")

	res = new(pb.ListCompaniesResponse)
	res.Success = false

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	comps, err := s.logic.ListCompanies(ctx, req.Filter)
	if err != nil {
		log.Errorw("failed to list companies", "err", err)
		return nil, err
	}

	res.Success = true
	res.Companies = &dto.CompanyList{
		Items:      comps.Items,
		Page:       int32(comps.Page),
		PerPage:    int32(comps.PerPage),
		TotalPages: int32(comps.TotalPages),
		TotalItems: int32(comps.TotalItems),
	}
	return
}

func (s *ServerGRPC) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (res *pb.GetCompanyResponse, err error) {
	log := s.log.Named("GetCompany").With(
		zap.String("company_id", req.GetId()),
	)
	log.Debug("GetCompany called")

	res = new(pb.GetCompanyResponse)
	res.Success = false

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	comp, err := s.logic.GetCompany(ctx, req.GetId())
	if err != nil {
		log.Errorw("failed to get company", "err", err)
		return nil, err
	}

	res.Success = true
	res.Company = comp
	return
}
