package participant

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/timestamp"
)

type Participant struct {
	ID        int32               `db:"id"`
	UserId    string              `db:"user_id"`
	CompanyId string              `db:"company_id"`
	RoleId    string              `db:"role_id"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}
type Role struct {
	ID          uuid.UUID           `db:"id"`
	Name        string              `db:"name"`
	Permissions []string            `db:"permissions"`
	CreatedAt   timestamp.Timestamp `db:"created_at"`
	UpdatedAt   timestamp.Timestamp `db:"updated_at"`
}
type ParticipantInfo struct {
	CompanyID   uuid.UUID
	CompanyName string
	Role        Role
	JoinedAt    timestamp.Timestamp
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
func (p *ParticipantInfo) DTO() *dto.ParticipantInfo {
	return &dto.ParticipantInfo{
		CompanyId:   p.CompanyID.String(),
		CompanyName: p.CompanyName,
		Role: &dto.Role{
			Id:          p.Role.ID.String(),
			Name:        p.Role.Name,
			Permissions: p.Role.Permissions,
			CreatedAt:   timestamppb.New(p.Role.CreatedAt.Time),
			UpdatedAt:   timestamppb.New(p.Role.UpdatedAt.Time),
		},
		JoinedAt: timestamppb.New(p.JoinedAt.Time),
	}
}
func ParticipantFromDTO(participant *dto.Participant) Participant {
	return Participant{
		ID:        participant.Id,
		UserId:    participant.UserId,
		RoleId:    participant.RoleId,
		CompanyId: participant.CompanyId,
		CreatedAt: timestamp.NewTimestamp(participant.CreatedAt.AsTime()),
		UpdatedAt: timestamp.NewTimestamp(participant.UpdatedAt.AsTime()),
	}
}
func ParticipantInfoFromDTO(info *dto.ParticipantInfo) ParticipantInfo {
	return ParticipantInfo{
		CompanyID:   uuid.MustParse((info.CompanyId)),
		CompanyName: info.CompanyName,
		Role: Role{
			ID:          uuid.MustParse(info.Role.Id),
			Name:        info.Role.Name,
			Permissions: info.Role.Permissions,
			CreatedAt:   timestamp.NewTimestamp(info.Role.CreatedAt.AsTime()),
			UpdatedAt:   timestamp.NewTimestamp(info.Role.UpdatedAt.AsTime()),
		},
		JoinedAt: timestamp.NewTimestamp(info.JoinedAt.AsTime()),
	}
}
