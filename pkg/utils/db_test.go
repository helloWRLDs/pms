package utils

import (
	"testing"
	"time"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func Test_GetColumns(t *testing.T) {
	user := User{ID: "1", Name: "Bob", CreatedAt: time.Now()}
	t.Log(GetColumns(user))
	t.Log(GetArguments(user)...)
}

type TestDBStruct struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Age       int    `db:"age"`
	Ignored   string `db:"-"`
	NoTag     string
	ZeroValue string `db:"zero_value"`
}

func TestGetColumns(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{
			name: "struct with non-zero values",
			input: TestDBStruct{
				ID:   "123",
				Name: "John",
				Age:  30,
			},
			expected: []string{"id", "name", "age"},
		},
		{
			name: "struct with zero values",
			input: TestDBStruct{
				ID:   "123",
				Name: "",
				Age:  0,
			},
			expected: []string{"id"},
		},
		{
			name: "pointer to struct",
			input: &TestDBStruct{
				ID:   "123",
				Name: "John",
				Age:  30,
			},
			expected: []string{"id", "name", "age"},
		},
		{
			name:     "non-struct type",
			input:    "not a struct",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetColumns(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("GetColumns() length = %v, want %v", len(got), len(tt.expected))
			}
			for i, v := range tt.expected {
				if got[i] != v {
					t.Errorf("GetColumns()[%d] = %v, want %v", i, got[i], v)
				}
			}
		})
	}
}

func TestGetArguments(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []interface{}
	}{
		{
			name: "struct with non-zero values",
			input: TestDBStruct{
				ID:   "123",
				Name: "John",
				Age:  30,
			},
			expected: []interface{}{"123", "John", 30},
		},
		{
			name: "struct with zero values",
			input: TestDBStruct{
				ID:   "123",
				Name: "",
				Age:  0,
			},
			expected: []interface{}{"123"},
		},
		{
			name: "pointer to struct",
			input: &TestDBStruct{
				ID:   "123",
				Name: "John",
				Age:  30,
			},
			expected: []interface{}{"123", "John", 30},
		},
		{
			name:     "non-struct type",
			input:    "not a struct",
			expected: []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetArguments(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("GetArguments() length = %v, want %v", len(got), len(tt.expected))
			}
			for i, v := range tt.expected {
				if got[i] != v {
					t.Errorf("GetArguments()[%d] = %v, want %v", i, got[i], v)
				}
			}
		})
	}
}
