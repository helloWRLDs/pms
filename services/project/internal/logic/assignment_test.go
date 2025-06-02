package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnassignTask(t *testing.T) {
	tests := []struct {
		name    string
		taskID  string
		userID  string
		wantErr bool
	}{
		{
			name:    "unassign existing task",
			taskID:  "02aa2692-3ea4-4bb9-8185-8cf7fd0dd466",
			userID:  "f3cef382-559d-4248-9b02-9c0038725ab7",
			wantErr: false,
		},
		{
			name:    "unassign non-existent task",
			taskID:  "non-existent-task",
			userID:  "f3cef382-559d-4248-9b02-9c0038725ab7",
			wantErr: false,
		},
		{
			name:    "unassign task from non-existent user",
			taskID:  "02aa2692-3ea4-4bb9-8185-8cf7fd0dd466",
			userID:  "non-existent-user",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.UnassignTask(context.Background(), tt.userID, tt.taskID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
