package data

import (
	"context"
	"testing"
	"time"

	"pms.pkg/type/list"
	"pms.pkg/type/timestamp"
)

func Test_Time(t *testing.T) {
	unix := 1746874231680
	ts := timestamp.NewTimestamp(time.UnixMilli(int64(unix)))
	t.Log(ts)
	t.Log(time.Now().Unix())
}

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

func Test_CountTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count := repo.Task.Count(ctx, list.Filters{
		Fields: map[string]string{
			"t.project_id": "1",
		},
	})
	t.Log(count)
}
