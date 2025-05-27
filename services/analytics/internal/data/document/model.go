package documentdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
)

type Document struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Body      *[]byte   `db:"body"`
	ProjectID string    `db:"project_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (d Document) DTO() *dto.Document {
	var body []byte
	if d.Body != nil {
		body = *d.Body
	}
	return &dto.Document{
		Id:        d.ID,
		Title:     d.Title,
		Body:      body,
		ProjectId: d.ProjectID,
		CreatedAt: timestamppb.New(d.CreatedAt),
	}
}

func Entity(dto *dto.Document) *Document {
	var body *[]byte
	if dto.Body != nil {
		bodyBytes := dto.Body
		body = &bodyBytes
	}
	return &Document{
		ID:        dto.Id,
		Title:     dto.Title,
		Body:      body,
		ProjectID: dto.ProjectId,
		CreatedAt: dto.CreatedAt.AsTime(),
	}
}
