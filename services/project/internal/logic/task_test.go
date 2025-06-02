package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name     string
		creation *dto.TaskCreation
		wantErr  bool
	}{
		{
			name: "create valid task",
			creation: &dto.TaskCreation{
				Title:      "Complete dark theme for home page",
				Body:       "Use Tailwindcss",
				Status:     string(consts.TASK_STATUS_IN_PROGRESS),
				ProjectId:  "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				Priority:   int32(consts.TASK_PRIORITY_HIGH),
				AssigneeId: "f3cef382-559d-4248-9b02-9c0038725ab7",
				DueDate:    timestamppb.New(time.Now().Add(72 * time.Hour)),
			},
			wantErr: false,
		},
		{
			name: "create task with empty title",
			creation: &dto.TaskCreation{
				Title:      "",
				Body:       "Use Tailwindcss",
				Status:     string(consts.TASK_STATUS_IN_PROGRESS),
				ProjectId:  "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				Priority:   int32(consts.TASK_PRIORITY_HIGH),
				AssigneeId: "f3cef382-559d-4248-9b02-9c0038725ab7",
				DueDate:    timestamppb.New(time.Now().Add(72 * time.Hour)),
			},
			wantErr: true,
		},
		{
			name: "create task with invalid project",
			creation: &dto.TaskCreation{
				Title:      "Complete dark theme for home page",
				Body:       "Use Tailwindcss",
				Status:     string(consts.TASK_STATUS_IN_PROGRESS),
				ProjectId:  "invalid-project-id",
				Priority:   int32(consts.TASK_PRIORITY_HIGH),
				AssigneeId: "f3cef382-559d-4248-9b02-9c0038725ab7",
				DueDate:    timestamppb.New(time.Now().Add(72 * time.Hour)),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := logic.CreateTask(context.Background(), tt.creation)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, id)
		})
	}
}

func TestGetTask(t *testing.T) {
	tests := []struct {
		name    string
		taskID  string
		wantErr bool
	}{
		{
			name:    "get existing task",
			taskID:  "712b8a41-2351-4286-ad03-086eaee4c417",
			wantErr: false,
		},
		{
			name:    "get non-existent task",
			taskID:  "non-existent-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := logic.GetTask(context.Background(), tt.taskID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, task)
			assert.Equal(t, tt.taskID, task.Id)
		})
	}
}

func TestListTasks(t *testing.T) {
	tests := []struct {
		name      string
		filter    *dto.TaskFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list all tasks",
			filter: &dto.TaskFilter{
				Page:    1,
				PerPage: 10,
			},
			wantCount: 10,
			wantErr:   false,
		},
		{
			name: "list tasks by project",
			filter: &dto.TaskFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
			},
			wantCount: 5,
			wantErr:   false,
		},
		{
			name: "list tasks by status",
			filter: &dto.TaskFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				Status:    string(consts.TASK_STATUS_DONE),
			},
			wantCount: 2,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks, err := logic.ListTasks(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, tasks.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, tasks.Page)
			assert.Equal(t, tt.filter.PerPage, tasks.PerPage)
		})
	}
}

func TestAssignTask(t *testing.T) {
	tests := []struct {
		name    string
		taskID  string
		userID  string
		wantErr bool
	}{
		{
			name:    "assign task to valid user",
			taskID:  "02aa2692-3ea4-4bb9-8185-8cf7fd0dd466",
			userID:  "f3cef382-559d-4248-9b02-9c0038725ab7",
			wantErr: false,
		},
		{
			name:    "assign non-existent task",
			taskID:  "non-existent-task",
			userID:  "f3cef382-559d-4248-9b02-9c0038725ab7",
			wantErr: true,
		},
		{
			name:    "assign task to non-existent user",
			taskID:  "02aa2692-3ea4-4bb9-8185-8cf7fd0dd466",
			userID:  "non-existent-user",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.AssignTask(context.Background(), tt.userID, tt.taskID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name    string
		taskID  string
		updates *dto.Task
		wantErr bool
	}{
		{
			name:   "update task title",
			taskID: "712b8a41-2351-4286-ad03-086eaee4c417",
			updates: &dto.Task{
				Id:        "712b8a41-2351-4286-ad03-086eaee4c417",
				Title:     "Updated Task Title",
				Body:      "Original Body",
				Status:    string(consts.TASK_STATUS_IN_PROGRESS),
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				Priority:  int32(consts.TASK_PRIORITY_HIGH),
				DueDate:   timestamppb.New(time.Now().Add(72 * time.Hour)),
			},
			wantErr: false,
		},
		{
			name:   "update task status",
			taskID: "712b8a41-2351-4286-ad03-086eaee4c417",
			updates: &dto.Task{
				Id:        "712b8a41-2351-4286-ad03-086eaee4c417",
				Title:     "Original Title",
				Body:      "Original Body",
				Status:    string(consts.TASK_STATUS_DONE),
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				Priority:  int32(consts.TASK_PRIORITY_HIGH),
				DueDate:   timestamppb.New(time.Now().Add(72 * time.Hour)),
			},
			wantErr: false,
		},
		{
			name:   "update non-existent task",
			taskID: "non-existent-id",
			updates: &dto.Task{
				Id:        "non-existent-id",
				Title:     "Updated Task Title",
				Body:      "Original Body",
				Status:    string(consts.TASK_STATUS_IN_PROGRESS),
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				Priority:  int32(consts.TASK_PRIORITY_HIGH),
				DueDate:   timestamppb.New(time.Now().Add(72 * time.Hour)),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.UpdateTask(context.Background(), tt.taskID, tt.updates)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			updated, err := logic.GetTask(context.Background(), tt.taskID)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, tt.updates.Title, updated.Title)
			assert.Equal(t, tt.updates.Status, updated.Status)
		})
	}
}

func TestListTasksWithFilters(t *testing.T) {
	tests := []struct {
		name      string
		filter    *dto.TaskFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list tasks by priority",
			filter: &dto.TaskFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				Priority:  int32(consts.TASK_PRIORITY_HIGH),
			},
			wantCount: 3,
			wantErr:   false,
		},
		{
			name: "list tasks by assignee",
			filter: &dto.TaskFilter{
				Page:       1,
				PerPage:    10,
				ProjectId:  "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				AssigneeId: "f3cef382-559d-4248-9b02-9c0038725ab7",
			},
			wantCount: 4,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks, err := logic.ListTasks(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, tasks.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, tasks.Page)
			assert.Equal(t, tt.filter.PerPage, tasks.PerPage)
		})
	}
}
