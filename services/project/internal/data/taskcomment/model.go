package taskcommentdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
)

type TaskComment struct {
	ID        string    `db:"id"`
	Body      string    `db:"body"`
	TaskID    string    `db:"task_id"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (ta TaskComment) DTO() *dto.TaskComment {
	return &dto.TaskComment{
		Id:     ta.ID,
		Body:   ta.Body,
		TaskId: ta.TaskID,
		User: &dto.User{
			Id: ta.UserID,
		},
		CreatedAt: timestamppb.New(ta.CreatedAt),
	}
}

func Entity(dto *dto.TaskComment) TaskComment {
	return TaskComment{
		ID:        dto.Id,
		Body:      dto.Body,
		TaskID:    dto.TaskId,
		UserID:    dto.User.Id,
		CreatedAt: dto.CreatedAt.AsTime(),
	}
}
