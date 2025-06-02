package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.api-gateway/internal/config"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
	"pms.pkg/transport/grpc/dto"
)

func TestIntegration_Task(t *testing.T) {
	cfg := config.Config{
		Auth: configgrpc.ClientConfig{
			Host: "localhost:50051",
		},
		Project: configgrpc.ClientConfig{
			Host: "localhost:50052",
		},
	}
	log := logger.Log

	l := &Logic{
		Config: cfg,
		log:    log,
	}

	t.Run("Create and Get Task", func(t *testing.T) {
		creation := &dto.TaskCreation{
			Title:     "Test Task",
			Body:      "Test Task Description",
			Status:    "TODO",
			Priority:  3,
			ProjectId: "test-project-1",
			Type:      "TASK",
		}
		err := l.CreateTask(context.Background(), creation)
		require.NoError(t, err)

		task, err := l.GetTask(context.Background(), "test-task-1")
		require.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, "Test Task", task.Title)
		assert.Equal(t, "Test Task Description", task.Body)
		assert.Equal(t, "TODO", task.Status)
		assert.Equal(t, int32(3), task.Priority)
		assert.Equal(t, "test-project-1", task.ProjectId)
		assert.Equal(t, "TASK", task.Type)
	})

	t.Run("List Tasks", func(t *testing.T) {
		filter := &dto.TaskFilter{
			Page:      1,
			PerPage:   10,
			ProjectId: "test-project-1", // Replace with actual project ID
		}
		tasks, err := l.ListTasks(context.Background(), filter)
		require.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.GreaterOrEqual(t, tasks.TotalItems, int32(0))
		assert.GreaterOrEqual(t, tasks.TotalPages, int32(0))

		for _, task := range tasks.Items {
			assert.NotEmpty(t, task.Id)
			assert.NotEmpty(t, task.Title)
			assert.NotEmpty(t, task.Status)
			assert.NotEmpty(t, task.ProjectId)
		}
	})

	t.Run("Update Task", func(t *testing.T) {
		update := &dto.Task{
			Id:        "test-task-1",
			Title:     "Updated Task Title",
			Body:      "Updated Task Description",
			Status:    "IN_PROGRESS",
			Priority:  4,
			ProjectId: "test-project-1", // Replace with actual project ID
			Type:      "TASK",
			DueDate:   timestamppb.New(time.Now().Add(24 * time.Hour)),
		}
		err := l.UpdateTask(context.Background(), update.Id, update)
		require.NoError(t, err)

		task, err := l.GetTask(context.Background(), update.Id)
		require.NoError(t, err)
		assert.Equal(t, "Updated Task Title", task.Title)
		assert.Equal(t, "Updated Task Description", task.Body)
		assert.Equal(t, "IN_PROGRESS", task.Status)
		assert.Equal(t, int32(4), task.Priority)
	})

	t.Run("Task Assignment", func(t *testing.T) {

		err := l.TaskAssign(context.Background(), "test-task-1", "test-user-1") // Replace with actual IDs
		require.NoError(t, err)

		err = l.TaskUnassign(context.Background(), "test-task-1", "test-user-1") // Replace with actual IDs
		require.NoError(t, err)
	})

	t.Run("Get Task with Invalid ID", func(t *testing.T) {
		task, err := l.GetTask(context.Background(), "non-existent-task")
		assert.Error(t, err)
		assert.Nil(t, task)
	})

	t.Run("Create Task with Invalid Data", func(t *testing.T) {
		creation := &dto.TaskCreation{
			Title: "",
		}
		err := l.CreateTask(context.Background(), creation)
		assert.Error(t, err)
	})

	t.Run("Update Task with Invalid ID", func(t *testing.T) {
		update := &dto.Task{
			Id: "non-existent-task",
		}
		err := l.UpdateTask(context.Background(), update.Id, update)
		assert.Error(t, err)
	})

	t.Run("Task Assignment with Invalid IDs", func(t *testing.T) {
		err := l.TaskAssign(context.Background(), "invalid-task", "invalid-user")
		assert.Error(t, err)
	})
}
