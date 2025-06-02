package consts

import (
	"testing"
)

func TestTaskType_String(t *testing.T) {
	tests := []struct {
		name     string
		taskType TaskType
		expected string
	}{
		{
			name:     "bug type",
			taskType: TaskTypeBug,
			expected: "bug",
		},
		{
			name:     "story type",
			taskType: TaskTypeStory,
			expected: "story",
		},
		{
			name:     "sub task type",
			taskType: TaskTypeSubTask,
			expected: "sub_task",
		},
		{
			name:     "feature type",
			taskType: TaskTypeFeature,
			expected: "feature",
		},
		{
			name:     "chore type",
			taskType: TaskTypeChore,
			expected: "chore",
		},
		{
			name:     "refactor type",
			taskType: TaskTypeRefactor,
			expected: "refactor",
		},
		{
			name:     "test type",
			taskType: TaskTypeTest,
			expected: "test",
		},
		{
			name:     "documentation type",
			taskType: TaskTypeDocumentation,
			expected: "documentation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.taskType.String(); got != tt.expected {
				t.Errorf("TaskType.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}
