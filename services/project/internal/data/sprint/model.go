package sprintdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type Sprint struct {
	ID          string     `db:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	StartDate   time.Time  `db:"start_date"`
	EndDate     time.Time  `db:"end_date"`
	ProjectID   string     `db:"project_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

func (s *Sprint) DTO() *dto.Sprint {
	return &dto.Sprint{
		Id:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		StartDate:   timestamppb.New(s.StartDate),
		EndDate:     timestamppb.New(s.EndDate),
		ProjectId:   s.ProjectID,
		CreatedAt:   timestamppb.New(s.CreatedAt),
		UpdatedAt:   timestamppb.New(utils.Value(s.UpdatedAt)),
	}
}

func Entity(dto *dto.Sprint) *Sprint {
	return &Sprint{
		ID:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
		StartDate:   dto.StartDate.AsTime(),
		EndDate:     dto.EndDate.AsTime(),
		ProjectID:   dto.ProjectId,
		CreatedAt:   dto.CreatedAt.AsTime(),
		UpdatedAt:   utils.Ptr(dto.UpdatedAt.AsTime()),
	}
}
