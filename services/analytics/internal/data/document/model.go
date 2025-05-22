package documentdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
)

type Document struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Body      []byte    `db:"body"`
	ProjectID string    `db:"project_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (d Document) DTO() *dto.Document {
	return &dto.Document{
		Id:        d.ID,
		Title:     d.Title,
		Body:      d.Body,
		ProjectId: d.ProjectID,
		CreatedAt: timestamppb.New(d.CreatedAt),
	}
}

func Entity(dto *dto.Document) *Document {
	return &Document{
		ID:        dto.Id,
		Title:     dto.Title,
		Body:      dto.Body,
		ProjectID: dto.ProjectId,
		CreatedAt: dto.CreatedAt.AsTime(),
	}
}
