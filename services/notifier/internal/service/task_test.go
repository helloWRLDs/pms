package service

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"pms.notifier/internal/modules/email/render"
// )

// type MockEmail struct {
// 	sendFunc func(data *render.EmailData, to string) error
// }

// func (m *MockEmail) Send(data *render.EmailData, to string) error {
// 	return m.sendFunc(data, to)
// }

// func TestNotifyTaskAssignment(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		assigneeName  string
// 		assigneeEmail string
// 		taskName      string
// 		taskId        string
// 		projectName   string
// 		wantErr       bool
// 	}{
// 		{
// 			name:          "successful notification",
// 			assigneeName:  "John Doe",
// 			assigneeEmail: "john@example.com",
// 			taskName:      "Implement Login",
// 			taskId:        "task-123",
// 			projectName:   "Project X",

// 			wantErr: false,
// 		},
// 		{
// 			name:          "email send failure",
// 			assigneeName:  "Jane Doe",
// 			assigneeEmail: "jane@example.com",
// 			taskName:      "Fix Bug",
// 			taskId:        "task-456",
// 			projectName:   "Project Y",
// 			mockSend: func(data *render.EmailData, to string) error {
// 				return assert.AnError
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name:          "invalid email",
// 			assigneeName:  "Invalid User",
// 			assigneeEmail: "invalid-email",
// 			taskName:      "Test Task",
// 			taskId:        "task-789",
// 			projectName:   "Project Z",
// 			mockSend: func(data *render.EmailData, to string) error {
// 				return nil
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockEmail := &MockEmail{sendFunc: tt.mockSend}
// 			service := &NotifierService{Email: mockEmail}

// 			err := service.NotifyTaskAssignment(
// 				context.Background(),
// 				tt.assigneeName,
// 				tt.assigneeEmail,
// 				tt.taskName,
// 				tt.taskId,
// 				tt.projectName,
// 			)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestNotifyTaskStatusChange(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		assigneeName  string
// 		assigneeEmail string
// 		taskName      string
// 		taskId        string
// 		projectName   string
// 		oldStatus     string
// 		newStatus     string
// 		mockSend      func(data *render.EmailData, to string) error
// 		wantErr       bool
// 	}{
// 		{
// 			name:          "successful status change notification",
// 			assigneeName:  "John Doe",
// 			assigneeEmail: "john@example.com",
// 			taskName:      "Implement Login",
// 			taskId:        "task-123",
// 			projectName:   "Project X",
// 			oldStatus:     "In Progress",
// 			newStatus:     "Completed",
// 			mockSend: func(data *render.EmailData, to string) error {
// 				assert.Equal(t, "john@example.com", to)
// 				assert.Equal(t, "TaskFlow", data.CompanyName)
// 				return nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:          "email send failure",
// 			assigneeName:  "Jane Doe",
// 			assigneeEmail: "jane@example.com",
// 			taskName:      "Fix Bug",
// 			taskId:        "task-456",
// 			projectName:   "Project Y",
// 			oldStatus:     "To Do",
// 			newStatus:     "In Progress",
// 			mockSend: func(data *render.EmailData, to string) error {
// 				return assert.AnError
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockEmail := &MockEmail{sendFunc: tt.mockSend}
// 			service := &NotifierService{Email: mockEmail}

// 			err := service.NotifyTaskStatusChange(
// 				context.Background(),
// 				tt.assigneeName,
// 				tt.assigneeEmail,
// 				tt.taskName,
// 				tt.taskId,
// 				tt.projectName,
// 				tt.oldStatus,
// 				tt.newStatus,
// 			)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestNotifyTaskComment(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		assigneeName  string
// 		assigneeEmail string
// 		taskName      string
// 		taskId        string
// 		projectName   string
// 		commentAuthor string
// 		commentText   string
// 		mockSend      func(data *render.EmailData, to string) error
// 		wantErr       bool
// 	}{
// 		{
// 			name:          "successful comment notification",
// 			assigneeName:  "John Doe",
// 			assigneeEmail: "john@example.com",
// 			taskName:      "Implement Login",
// 			taskId:        "task-123",
// 			projectName:   "Project X",
// 			commentAuthor: "Jane Smith",
// 			commentText:   "Please review the changes",
// 			mockSend: func(data *render.EmailData, to string) error {
// 				assert.Equal(t, "john@example.com", to)
// 				assert.Equal(t, "TaskFlow", data.CompanyName)
// 				return nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:          "email send failure",
// 			assigneeName:  "Jane Doe",
// 			assigneeEmail: "jane@example.com",
// 			taskName:      "Fix Bug",
// 			taskId:        "task-456",
// 			projectName:   "Project Y",
// 			commentAuthor: "John Smith",
// 			commentText:   "I've fixed the issue",
// 			mockSend: func(data *render.EmailData, to string) error {
// 				return assert.AnError
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockEmail := &MockEmail{sendFunc: tt.mockSend}
// 			service := &NotifierService{Email: mockEmail}

// 			err := service.NotifyTaskComment(
// 				context.Background(),
// 				tt.assigneeName,
// 				tt.assigneeEmail,
// 				tt.taskName,
// 				tt.taskId,
// 				tt.projectName,
// 				tt.commentAuthor,
// 				tt.commentText,
// 			)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestEmailTemplateRendering(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		content interface{}
// 		wantErr bool
// 	}{
// 		{
// 			name: "valid task assignment content",
// 			content: render.TaskAssignmentContent{
// 				AssigneeName: "John Doe",
// 				TaskName:     "Implement Login",
// 				TaskId:       "task-123",
// 				ProjectName:  "Project X",
// 				CompanyName:  "TaskFlow",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "valid task status change content",
// 			content: render.TaskStatusChangeContent{
// 				AssigneeName: "John Doe",
// 				TaskName:     "Implement Login",
// 				TaskId:       "task-123",
// 				ProjectName:  "Project X",
// 				OldStatus:    "In Progress",
// 				NewStatus:    "Completed",
// 				CompanyName:  "TaskFlow",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "valid task comment content",
// 			content: render.TaskCommentContent{
// 				AssigneeName:  "John Doe",
// 				TaskName:      "Implement Login",
// 				TaskId:        "task-123",
// 				ProjectName:   "Project X",
// 				CommentAuthor: "Jane Smith",
// 				CommentText:   "Please review the changes",
// 				CompanyName:   "TaskFlow",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:    "empty content",
// 			content: nil,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			data, err := render.Render(tt.content)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Nil(t, data)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, data)
// 				assert.Equal(t, "TaskFlow", data.CompanyName)
// 			}
// 		})
// 	}
// }
