package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/timestamp"
	"pms.pkg/utils"
)

type Task struct {
	ID        string              `db:"id"`
	Title     string              `db:"title"`
	Body      string              `db:"body"`
	Status    consts.TaskStatus   `db:"status"`
	Priority  *int                `db:"priority"`
	ProjectID string              `db:"project_id"`
	SprintID  *string             `db:"sprint_id"`
	DueDate   timestamp.Timestamp `db:"due_date"`
	Created   timestamp.Timestamp `db:"created_at"`
	Updated   timestamp.Timestamp `db:"updated_at"`
}

func (t Task) DTO() *dto.Task {
	return &dto.Task{
		Id:        t.ID,
		Title:     t.Title,
		Body:      t.Body,
		Status:    string(t.Status),
		SprintId:  utils.Value(t.SprintID),
		ProjectId: t.ProjectID,
		Priority:  int32(utils.Value(t.Priority)),
		CreatedAt: timestamppb.New(t.Created.Time),
		UpdatedAt: timestamppb.New(t.Updated.Time),
		DueDate:   timestamppb.New(t.DueDate.Time),
	}
}
