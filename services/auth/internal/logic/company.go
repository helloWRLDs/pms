package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	companydata "pms.auth/internal/data/company"
	participantdata "pms.auth/internal/data/participant"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
)

func (l *Logic) UpdateCompany(ctx context.Context, companyID string, updatedCompany *dto.Company) (err error) {
	log := l.log.With(
		zap.String("func", "UpdateCompany"),
		zap.Any("updated_company", updatedCompany),
	)
	log.Debug("UpdateCompany called")

	if exist := l.Repo.Company.Exists(ctx, "id", companyID); !exist {
		log.Error("company not found")
		return errs.ErrNotFound{
			Object: "company",
			Field:  "id",
			Value:  updatedCompany.Id,
		}
	}

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		log.Infow("err check", "isNil", err == nil, "err", err)
		l.Repo.EndTx(tx, err)
	}()

	company := companydata.Company{
		ID:          companyID,
		Name:        updatedCompany.Name,
		Codename:    updatedCompany.Codename,
		BIN:         updatedCompany.Bin,
		Address:     updatedCompany.Address,
		Description: updatedCompany.Description,
	}
	if err := l.Repo.Company.UpdateCompany(tx, companyID, company); err != nil {
		log.Errorw("failed to update company", "err", err)
		return err
	}
	return nil
}

func (l *Logic) CreateCompany(ctx context.Context, userID string, newCompany *dto.NewCompany) (created *dto.Company, err error) {
	log := l.log.With(
		zap.String("func", "CreateCompany"),
		zap.Any("new_company", newCompany),
	)
	log.Debug("CreateCompany called")

	if exist := l.Repo.User.Exists(ctx, "id", userID); !exist {
		log.Error("user not found")
		return nil, errs.ErrNotFound{
			Object: "user",
			Field:  "id",
			Value:  userID,
		}
	}

	company := companydata.Company{
		ID:          uuid.NewString(),
		Name:        newCompany.Name,
		Codename:    newCompany.Codename,
		BIN:         newCompany.Bin,
		Address:     newCompany.Address,
		Description: newCompany.Description,
	}

	participant := participantdata.Participant{
		ID:        uuid.NewString(),
		UserID:    userID,
		CompanyID: company.ID,
		Role:      "admin",
	}

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		log.Infow("err check", "isNil", err == nil, "err", err)
		l.Repo.EndTx(tx, err)
	}()

	if err := l.Repo.Company.Create(tx, company); err != nil {
		log.Errorw("failed to create company", "err", err)
		return nil, err
	}

	if err := l.Repo.Participant.Create(tx, participant); err != nil {
		log.Errorw("failed to create company", "err", err)
		return nil, err
	}
	log.Infow("company created", "admin_id", userID)

	created = new(dto.Company)
	created.Id = company.ID
	created.Name = company.Name
	created.Codename = company.Codename
	created.Bin = company.BIN
	created.Address = company.Address
	created.Description = company.Description
	created, err = l.GetCompany(tx, company.ID)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (l *Logic) ListCompanies(ctx context.Context, filter *dto.CompanyFilter) (list.List[*dto.Company], error) {
	log := l.log.With(
		zap.String("func", "ListCompanies"),
		zap.Any("filter", filter),
	)
	log.Info("ListCompanies called")

	comps, err := l.Repo.Company.List(ctx, filter)
	if err != nil {
		log.Errorw("failed to list companies", "err", err)
		return list.List[*dto.Company]{}, err
	}
	log.Infow("companies found", "companies", comps)
	result := list.List[*dto.Company]{
		Pagination: list.Pagination{
			Page:       comps.Page,
			PerPage:    comps.PerPage,
			TotalItems: comps.TotalItems,
			TotalPages: comps.TotalPages,
		},
	}
	for _, comp := range comps.Items {
		result.Items = append(result.Items, func() (c *dto.Company) {
			c = comp.DTO()
			c.PeopleCount = l.Repo.Company.Count(ctx, &dto.CompanyFilter{
				CompanyId: comp.ID,
			})
			return
		}())
	}

	return result, nil
}

func (l *Logic) GetCompany(ctx context.Context, id string) (*dto.Company, error) {
	log := l.log.With(
		zap.String("func", "GetCompany"),
		zap.String("id", id),
	)
	log.Debug("GetCompany called")

	company := new(dto.Company)

	comp, err := l.Repo.Company.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	company.Id = comp.ID
	company.Name = comp.Name
	company.Codename = comp.Codename
	company.Bin = comp.BIN
	company.Address = comp.Address
	company.Description = comp.Description

	company.PeopleCount = l.Repo.Company.Count(ctx, &dto.CompanyFilter{
		CompanyId: id,
	})

	return company, nil
}
