package render

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render(t *testing.T) {
	greet := GreetContent{
		"Bob", "AITU",
	}
	t.Logf("%#v", mqTable(greet))
}

func mqTable(data interface{}) (table map[string]interface{}) {
	table = make(map[string]interface{}, 0)
	v := reflect.ValueOf(data)

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if !value.CanInterface() {
			continue
		}
		table[field.Name] = value.Interface()
	}
	return
}

func TestRender(t *testing.T) {
	tests := []struct {
		name    string
		content TaskAssignmentContent
		wantErr bool
	}{
		{
			name: "valid task assignment content",
			content: TaskAssignmentContent{
				AssigneeName: "John Doe",
				TaskName:     "Implement Login Feature",
				TaskId:       "task-123",
				ProjectName:  "Project X",
				CompanyName:  "TaskFlow",
			},
			wantErr: false,
		},
		{
			name: "empty content",
			content: TaskAssignmentContent{
				AssigneeName: "",
				TaskName:     "",
				TaskId:       "",
				ProjectName:  "",
				CompanyName:  "TaskFlow",
			},
			wantErr: false,
		},
		{
			name: "missing company name",
			content: TaskAssignmentContent{
				AssigneeName: "John Doe",
				TaskName:     "Implement Login Feature",
				TaskId:       "task-123",
				ProjectName:  "Project X",
				CompanyName:  "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Render(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, data)

			content := string(data)
			if tt.content.AssigneeName != "" {
				assert.Contains(t, content, tt.content.AssigneeName)
			}
			if tt.content.TaskName != "" {
				assert.Contains(t, content, tt.content.TaskName)
			}
			if tt.content.TaskId != "" {
				assert.Contains(t, content, tt.content.TaskId)
			}
			if tt.content.ProjectName != "" {
				assert.Contains(t, content, tt.content.ProjectName)
			}
			if tt.content.CompanyName != "" {
				assert.Contains(t, content, tt.content.CompanyName)
			}
		})
	}
}

func TestRenderWithInvalidTemplate(t *testing.T) {

	content := TaskAssignmentContent{
		AssigneeName: "{{.InvalidField}}",
		TaskName:     "Test Task",
		TaskId:       "task-123",
		ProjectName:  "Project X",
		CompanyName:  "TaskFlow",
	}

	data, err := Render(content)
	assert.Error(t, err)
	assert.Empty(t, data)
}

func TestGreetContent(t *testing.T) {
	tests := []struct {
		name    string
		content GreetContent
		wantErr bool
	}{
		{
			name: "valid greet content",
			content: GreetContent{
				Name: "John Doe",
			},
			wantErr: false,
		},
		{
			name: "empty greet content",
			content: GreetContent{
				Name: "",
			},
			wantErr: false,
		},
		{
			name: "missing company name",
			content: GreetContent{
				Name: "John Doe",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Render(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, data)

			content := string(data)
			if tt.content.Name != "" {
				assert.Contains(t, content, tt.content.Name)
			}

		})
	}
}

func TestRenderWithSpecialCharacters(t *testing.T) {
	tests := []struct {
		name    string
		content TaskAssignmentContent
		wantErr bool
	}{
		{
			name: "content with HTML characters",
			content: TaskAssignmentContent{
				AssigneeName: "<script>alert('test')</script>",
				TaskName:     "Fix <b>HTML</b> rendering",
				TaskId:       "task-123",
				ProjectName:  "Project & Company",
				CompanyName:  "TaskFlow & Co.",
			},
			wantErr: false,
		},
		{
			name: "content with emojis",
			content: TaskAssignmentContent{
				AssigneeName: "John 😊 Doe",
				TaskName:     "Implement 🎨 UI",
				TaskId:       "task-123",
				ProjectName:  "Project 🚀",
				CompanyName:  "TaskFlow ⭐",
			},
			wantErr: false,
		},
		{
			name: "content with newlines",
			content: TaskAssignmentContent{
				AssigneeName: "John\nDoe",
				TaskName:     "Implement\nLogin",
				TaskId:       "task-123",
				ProjectName:  "Project\nX",
				CompanyName:  "TaskFlow\nInc",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Render(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, data)

			content := string(data)
			assert.Contains(t, content, tt.content.AssigneeName)
			assert.Contains(t, content, tt.content.TaskName)
			assert.Contains(t, content, tt.content.TaskId)
			assert.Contains(t, content, tt.content.ProjectName)
			assert.Contains(t, content, tt.content.CompanyName)
		})
	}
}

func TestRenderWithLongContent(t *testing.T) {
	longString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
		"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris " +
		"nisi ut aliquip ex ea commodo consequat."

	tests := []struct {
		name    string
		content TaskAssignmentContent
		wantErr bool
	}{
		{
			name: "long task name",
			content: TaskAssignmentContent{
				AssigneeName: "John Doe",
				TaskName:     longString,
				TaskId:       "task-123",
				ProjectName:  "Project X",
				CompanyName:  "TaskFlow",
			},
			wantErr: false,
		},
		{
			name: "long project name",
			content: TaskAssignmentContent{
				AssigneeName: "John Doe",
				TaskName:     "Implement Login",
				TaskId:       "task-123",
				ProjectName:  longString,
				CompanyName:  "TaskFlow",
			},
			wantErr: false,
		},
		{
			name: "long assignee name",
			content: TaskAssignmentContent{
				AssigneeName: longString,
				TaskName:     "Implement Login",
				TaskId:       "task-123",
				ProjectName:  "Project X",
				CompanyName:  "TaskFlow",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Render(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, data)

			content := string(data)
			assert.Contains(t, content, tt.content.AssigneeName)
			assert.Contains(t, content, tt.content.TaskName)
			assert.Contains(t, content, tt.content.TaskId)
			assert.Contains(t, content, tt.content.ProjectName)
			assert.Contains(t, content, tt.content.CompanyName)
		})
	}
}

func TestWelcomeLoginContent(t *testing.T) {
	tests := []struct {
		name    string
		content WelcomeLoginContent
		wantErr bool
	}{
		{
			name: "first login welcome content with full details",
			content: WelcomeLoginContent{
				Name:              "John Doe",
				CompanyName:       "TaskFlow Inc",
				CompanyAddress:    "123 Business St, Tech City, TC 12345",
				IsFirstLogin:      true,
				DashboardURL:      "https://taskflow.com/dashboard",
				SetupGuideURL:     "https://taskflow.com/setup",
				SupportEmail:      "support@taskflow.com",
				SupportPhone:      "+1-800-TASKFLOW",
				DocumentationURL:  "https://docs.taskflow.com",
				VideoTutorialsURL: "https://taskflow.com/tutorials",
				LoginDetails: &LoginDetails{
					Timestamp: "2024-01-15 10:30:00 UTC",
					Location:  "New York, NY, USA",
					Device:    "Chrome on Windows 11",
					IPAddress: "192.168.1.100",
				},
				SocialLinks: SocialLinks{
					Website:  "https://taskflow.com",
					Twitter:  "https://twitter.com/taskflow",
					LinkedIn: "https://linkedin.com/company/taskflow",
				},
			},
			wantErr: false,
		},
		{
			name: "returning user welcome content minimal",
			content: WelcomeLoginContent{
				Name:         "Jane Smith",
				CompanyName:  "TaskFlow Inc",
				IsFirstLogin: false,
				DashboardURL: "https://taskflow.com/dashboard",
				SupportEmail: "support@taskflow.com",
			},
			wantErr: false,
		},
		{
			name: "first login with no optional fields",
			content: WelcomeLoginContent{
				Name:         "Bob Johnson",
				CompanyName:  "Minimal Corp",
				IsFirstLogin: true,
				DashboardURL: "https://minimal.com/dashboard",
				SupportEmail: "help@minimal.com",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Render(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, data)

			content := string(data)

			// Verify required fields are present
			assert.Contains(t, content, tt.content.Name)
			assert.Contains(t, content, tt.content.CompanyName)
			assert.Contains(t, content, tt.content.DashboardURL)
			assert.Contains(t, content, tt.content.SupportEmail)

			// Verify subject line logic
			subject := tt.content.Subject()
			if tt.content.IsFirstLogin {
				assert.Contains(t, subject, "Let's Get Started!")
			} else {
				assert.Contains(t, subject, "Welcome back")
			}

			// Verify optional fields when present
			if tt.content.SetupGuideURL != "" {
				assert.Contains(t, content, tt.content.SetupGuideURL)
			}
			if tt.content.SupportPhone != "" {
				assert.Contains(t, content, tt.content.SupportPhone)
			}
			if tt.content.LoginDetails != nil {
				assert.Contains(t, content, tt.content.LoginDetails.Timestamp)
				assert.Contains(t, content, tt.content.LoginDetails.IPAddress)
			}
			if tt.content.SocialLinks.Website != "" {
				assert.Contains(t, content, tt.content.SocialLinks.Website)
			}
		})
	}
}
