package utils

import (
	"testing"
)

func TestPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "string pointer",
			input:    "test",
			expected: "test",
		},
		{
			name:     "int pointer",
			input:    42,
			expected: 42,
		},
		{
			name:     "zero value string",
			input:    "",
			expected: nil,
		},
		{
			name:     "zero value int",
			input:    0,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case string:
				result := Ptr(v)
				if tt.expected == nil {
					if result != nil {
						t.Errorf("Ptr() = %v, want nil", result)
					}
				} else if *result != tt.expected {
					t.Errorf("Ptr() = %v, want %v", *result, tt.expected)
				}
			case int:
				result := Ptr(v)
				if tt.expected == nil {
					if result != nil {
						t.Errorf("Ptr() = %v, want nil", result)
					}
				} else if *result != tt.expected {
					t.Errorf("Ptr() = %v, want %v", *result, tt.expected)
				}
			}
		})
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "string pointer",
			input:    stringPtr("test"),
			expected: "test",
		},
		{
			name:     "int pointer",
			input:    intPtr(42),
			expected: 42,
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case *string:
				result := Value(v)
				if result != tt.expected {
					t.Errorf("Value() = %v, want %v", result, tt.expected)
				}
			case *int:
				result := Value(v)
				if result != tt.expected {
					t.Errorf("Value() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
