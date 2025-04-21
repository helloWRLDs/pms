package entity

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/timestamp"
)

type Participant struct {
	ID        int32               `db:"id"`
	UserId    string              `db:"user_id"`
	CompanyId string              `db:"company_id"`
	RoleId    string              `db:"role"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}

func (p *Participant) DTO() *dto.Participant {
	return &dto.Participant{
		Id:        p.ID,
		UserId:    p.UserId,
		CompanyId: p.CompanyId,
		RoleId:    p.RoleId,
		CreatedAt: timestamppb.New(p.CreatedAt.Time),
		UpdatedAt: timestamppb.New(p.UpdatedAt.Time),
	}
}
