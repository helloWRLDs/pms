package service

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"pms.pkg/logger"
)

func Test_NotifyTaskAssignment(t *testing.T) {
	log := logger.Log.With(
		zap.String("test", "Test_NotifyTaskAssignment"),
	)
	log.Debug("test started")

	tests := []struct {
		name         string
		assigneeName string
		email        string
		taskName     string
		taskId       string
		projectName  string
		wantErr      bool
	}{
		{
			name:         "valid task assignment",
			assigneeName: "John Doe",
			email:        "kossinovviktor@gmail.com",
			taskName:     "Implement Login",
			taskId:       "task-123",
			projectName:  "Project X",
			wantErr:      false,
		},
		{
			name:         "invalid email",
			assigneeName: "Jane Doe",
			email:        "invalid-email",
			taskName:     "Fix Bugs",
			taskId:       "task-456",
			projectName:  "Project Y",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Notifier.NotifyTaskAssignment(
				context.Background(),
				tt.assigneeName,
				tt.email,
				tt.taskName,
				tt.taskId,
				tt.projectName,
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyTaskAssignment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				log.Infow("task assignment notification sent successfully",
					"assignee", tt.assigneeName,
					"email", tt.email,
					"task", tt.taskName,
				)
			}
		})
	}
}
