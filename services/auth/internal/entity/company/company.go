package company

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/timestamp"
)

type Company struct {
	ID        uuid.UUID           `db:"id"`
	Name      string              `db:"name"`
	Codename  string              `db:"codename"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}

func (c *Company) DTO() *dto.Company {
	return &dto.Company{
		Id:        c.ID.String(),
		Name:      c.Name,
		Codename:  c.Codename,
		CreatedAt: timestamppb.New(c.CreatedAt.Time),
		UpdatedAt: timestamppb.New(c.UpdatedAt.Time),
	}
}

func CompanyFromDTO(company *dto.Company) Company {
	return Company{
		ID:        uuid.MustParse(company.Id),
		Codename:  company.Codename,
		Name:      company.Name,
		CreatedAt: timestamp.NewTimestamp(company.CreatedAt.AsTime()),
		UpdatedAt: timestamp.NewTimestamp(company.UpdatedAt.AsTime()),
	}
}
