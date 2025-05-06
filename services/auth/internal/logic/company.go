package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
)

func (l *Logic) ListCompanies(ctx context.Context, userID string, filter list.Filters) (list.List[*dto.Company], error) {
	log := l.log.With(
		zap.String("func", "ListCompanies"),
		zap.String("user_id", userID),
		zap.Any("filter", filter),
	)
	log.Info("ListCompanies called")

	user, err := l.Repo.User.GetByID(ctx, userID)
	if err != nil {
		log.Errorw("failed to get user", "err", err)
		return list.List[*dto.Company]{}, err
	}
	if user.ID == uuid.Nil {
		log.Error("failed to get user")
		return list.List[*dto.Company]{}, errs.ErrNotFound{
			Object: "user",
			Field:  "id",
			Value:  userID,
		}
	}
	log.Infow("user found", "user", user)

	comps, err := l.Repo.Company.List(ctx, list.Filters{
		Pagination: list.Pagination{
			Page:    filter.Page,
			PerPage: filter.PerPage,
		},
		Fields: map[string]string{
			"p.user_id": userID,
		},
	})
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
		result.Items = append(result.Items, &dto.Company{
			Id:       comp.ID.String(),
			Name:     comp.Name,
			Codename: comp.Codename,
			PeopleCount: l.Repo.User.Count(ctx, list.Filters{
				Fields: map[string]string{
					"p.company_id": comp.ID.String(),
				},
				Pagination: list.Pagination{
					Page:    1,
					PerPage: 100000,
				},
			}),
			CreatedAt: timestamppb.New(comp.CreatedAt.Time),
			UpdatedAt: timestamppb.New(comp.UpdatedAt.Time),
		})
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

	company.Id = comp.ID.String()
	company.Name = comp.Name
	company.Codename = comp.Codename

	company.PeopleCount = l.Repo.User.Count(ctx, list.Filters{
		Fields: map[string]string{
			"p.company_id": id,
		},
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10000,
		},
	})

	return company, nil
}
