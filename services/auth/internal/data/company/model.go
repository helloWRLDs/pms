package companydata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type Company struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	Codename  string     `db:"codename"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func Entity(dto *dto.Company) *Company {
	return &Company{
		ID:        dto.Id,
		Name:      dto.Name,
		Codename:  dto.Codename,
		CreatedAt: dto.CreatedAt.AsTime(),
		UpdatedAt: utils.Ptr(dto.UpdatedAt.AsTime()),
	}
}

func (c *Company) DTO() *dto.Company {
	return &dto.Company{
		Id:        c.ID,
		Name:      c.Name,
		Codename:  c.Codename,
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(utils.Value(c.UpdatedAt)),
	}
}
