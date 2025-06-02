package utils

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
)

func TestCalculateTaskPoints(t *testing.T) {
	now := time.Now()
	baseTask := &dto.Task{
		Id:        "test-task",
		Type:      string(consts.TaskTypeFeature),
		Priority:  3,
		Status:    string(consts.TASK_STATUS_DONE),
		CreatedAt: timestamppb.New(now.Add(-24 * time.Hour)),
		UpdatedAt: timestamppb.New(now),
		DueDate:   timestamppb.New(now.Add(24 * time.Hour)),
	}

	tests := []struct {
		name     string
		task     *dto.Task
		expected int32
	}{
		{
			name:     "feature task with medium priority",
			task:     baseTask,
			expected: 25,
		},
		{
			name: "bug task with high priority",
			task: func() *dto.Task {
				task := *baseTask
				task.Type = string(consts.TaskTypeBug)
				task.Priority = 2
				return &task
			}(),
			expected: 19,
		},
		{
			name: "story task with highest priority",
			task: func() *dto.Task {
				task := *baseTask
				task.Type = string(consts.TaskTypeStory)
				task.Priority = 1
				return &task
			}(),
			expected: 30,
		},
		{
			name: "chore task with low priority",
			task: func() *dto.Task {
				task := *baseTask
				task.Type = string(consts.TaskTypeChore)
				task.Priority = 4
				return &task
			}(),
			expected: 6,
		},
		{
			name: "completed before due date",
			task: func() *dto.Task {
				task := *baseTask
				task.CreatedAt = timestamppb.New(now.Add(-48 * time.Hour))
				task.UpdatedAt = timestamppb.New(now.Add(-24 * time.Hour))
				task.DueDate = timestamppb.New(now)
				return &task
			}(),
			expected: 26,
		},
		{
			name: "completed after due date",
			task: func() *dto.Task {
				task := *baseTask
				task.CreatedAt = timestamppb.New(now.Add(-72 * time.Hour))
				task.UpdatedAt = timestamppb.New(now.Add(24 * time.Hour))
				task.DueDate = timestamppb.New(now)
				return &task
			}(),
			expected: 24,
		},
		{
			name: "quick completion bonus",
			task: func() *dto.Task {
				task := *baseTask
				task.CreatedAt = timestamppb.New(now.Add(-12 * time.Hour))
				task.UpdatedAt = timestamppb.New(now)
				return &task
			}(),
			expected: 30,
		},
		{
			name: "long completion penalty",
			task: func() *dto.Task {
				task := *baseTask
				task.CreatedAt = timestamppb.New(now.Add(-15 * 24 * time.Hour))
				task.UpdatedAt = timestamppb.New(now)
				return &task
			}(),
			expected: 22,
		},
		{
			name: "minimum points",
			task: func() *dto.Task {
				task := *baseTask
				task.Type = string(consts.TaskTypeChore)
				task.Priority = 5
				task.CreatedAt = timestamppb.New(now.Add(-30 * 24 * time.Hour))
				task.UpdatedAt = timestamppb.New(now.Add(30 * 24 * time.Hour))
				task.DueDate = timestamppb.New(now)
				return &task
			}(),
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTaskPoints(tt.task)
			if got != tt.expected {
				t.Errorf("CalculateTaskPoints() = %v, want %v", got, tt.expected)
			}
		})
	}
}
