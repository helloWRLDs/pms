package logic

import (
	"context"
	"testing"

	"pms.pkg/utils"
)

func TestGetUserTaskStats(t *testing.T) {
	companyID := "60cde332-ad5a-4aab-932b-81b5f16a61d2"
	stats, err := logic.GetUserTaskStats(context.Background(), companyID)
	if err != nil {
		t.Fatalf("failed to get user task stats: %v", err)
	}
	t.Log(utils.JSON(stats))
}
