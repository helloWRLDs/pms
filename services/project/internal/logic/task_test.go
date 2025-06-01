package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func Test_CreateTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creation := &dto.TaskCreation{
		Title:      "Complete dark theme for home page",
		Body:       "Use Tailwindcss",
		Status:     string(consts.TASK_STATUS_IN_PROGRESS),
		ProjectId:  "1",
		Priority:   int32(consts.TASK_PRIORITY_HIGH),
		AssigneeId: "user_2",
		DueDate:    timestamppb.New(time.Now().Add(72 * time.Hour)),
	}

	id, err := logic.CreateTask(ctx, creation)
	assert.NoError(t, err)
	t.Log(id)
}

func Test_GetTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	taskID := "712b8a41-2351-4286-ad03-086eaee4c417"

	task, err := logic.GetTask(ctx, taskID)
	assert.NoError(t, err)
	t.Log(utils.JSON(task))
}

func Test_ListTasks(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := &dto.TaskFilter{
		Page:      1,
		PerPage:   10,
		ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
		Status:    string(consts.TASK_STATUS_DONE),
	}

	list, err := logic.ListTasks(ctx, filter)
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}
func Test_ReassignTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	taskID := "02aa2692-3ea4-4bb9-8185-8cf7fd0dd466"
	userID := "f3cef382-559d-4248-9b02-9c0038725ab7"
	err := logic.AssignTask(ctx, userID, taskID)
	assert.NoError(t, err)
}
