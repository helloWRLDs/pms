package data

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
	taskdata "pms.project/internal/data/task"
)

func Test_GetTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	taskID := "3216d7eb-3695-46e0-a9da-056382f7f7b4"
	task, err := repo.Task.GetByID(ctx, taskID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(task)
}

func Test_ListTasks(t *testing.T) {
	projectID := "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1"
	l, err := repo.Task.List(context.Background(), &dto.TaskFilter{
		Page:      1,
		PerPage:   10,
		ProjectId: projectID,
	})

	assert.NoError(t, err)
	t.Log(utils.JSON(l))
}

func Test_CreateTask(t *testing.T) {
	newTask := taskdata.Task{
		Title:     "vite configuration",
		Body:      "Setup vite configuration and tailwind styles",
		Status:    consts.TASK_STATUS_CREATED,
		Priority:  utils.Ptr(3),
		ProjectID: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
	}
	err := repo.Task.Create(context.Background(), newTask)
	assert.NoError(t, err)
}

func Test_CountTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count := repo.Task.Count(ctx, &dto.TaskFilter{
		ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
	})
	t.Log(count)
}
