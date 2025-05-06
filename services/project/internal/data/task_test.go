package data

import (
	"context"
	"testing"
	"time"

	"pms.pkg/type/list"
)

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
