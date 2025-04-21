package grpchandler

import (
	"context"

	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/type/list"
)

func (s *ServerGRPC) ListCompanies(ctx context.Context, req *pb.ListCompaniesRequest) (res *pb.ListCompaniesResponse, err error) {
	log := s.log.With("func", "ListCompanies", "pkg", "grpchandler")
	log.Debug("ListCompanies called")

	res = new(pb.ListCompaniesResponse)
	res.Success = false

	defer func() {
		err = errs.WrapGRPC(err)
	}()
	log.Infof("userID: %s", req.UserId)

	filter := list.Filters{
		Pagination: list.Pagination{
			Page:    int(req.Page),
			PerPage: int(req.PerPage),
		},
	}

	comps, err := s.logic.ListCompanies(ctx, req.GetUserId(), filter)
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
	log := s.log.With("func", "RegisterUser", "pkg", "grpchandler")
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
