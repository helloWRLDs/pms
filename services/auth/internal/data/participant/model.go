package participantdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type Participant struct {
	ID        string     `db:"id"`
	UserID    string     `db:"user_id"`
	CompanyID string     `db:"company_id"`
	Role      string     `db:"role"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func (p *Participant) DTO() *dto.Participant {
	return &dto.Participant{
		Id:        p.ID,
		UserId:    p.UserID,
		CompanyId: p.CompanyID,
		Role:      p.Role,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(utils.Value(p.UpdatedAt)),
	}
}
