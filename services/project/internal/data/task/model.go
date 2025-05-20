package taskdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type Task struct {
	ID        string            `db:"id"`
	Title     string            `db:"title"`
	Body      string            `db:"body"`
	Status    consts.TaskStatus `db:"status"`
	Priority  *int              `db:"priority"`
	ProjectID string            `db:"project_id"`
	SprintID  *string           `db:"sprint_id"`
	Code      string            `db:"code"`
	DueDate   *time.Time        `db:"due_date"`
	Created   time.Time         `db:"created_at"`
	Updated   *time.Time        `db:"updated_at"`
}

func (t Task) DTO() *dto.Task {
	return &dto.Task{
		Id:        t.ID,
		Title:     t.Title,
		Body:      t.Body,
		Status:    string(t.Status),
		SprintId:  utils.Value(t.SprintID),
		Code:      t.Code,
		ProjectId: t.ProjectID,
		Priority:  int32(utils.Value(t.Priority)),
		CreatedAt: timestamppb.New(t.Created),
		UpdatedAt: timestamppb.New(utils.Value(t.Updated)),
		DueDate:   timestamppb.New(utils.Value(t.DueDate)),
	}
}

func Entity(dto *dto.Task) *Task {
	return &Task{
		ID:        dto.Id,
		Title:     dto.Title,
		Body:      dto.Body,
		Status:    consts.TaskStatus(dto.Status),
		Priority:  utils.Ptr(int(dto.Priority)),
		ProjectID: dto.ProjectId,
		Code:      dto.Code,
		SprintID:  utils.Ptr(dto.SprintId),
		Created:   dto.CreatedAt.AsTime(),
		Updated:   utils.Ptr(dto.UpdatedAt.AsTime()),
		DueDate:   utils.Ptr(dto.DueDate.AsTime()),
	}
}
