package utils

import (
	"testing"
)

func TestIf(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		a         interface{}
		b         interface{}
		expected  interface{}
	}{
		{
			name:      "true condition with strings",
			condition: true,
			a:         "yes",
			b:         "no",
			expected:  "yes",
		},
		{
			name:      "false condition with strings",
			condition: false,
			a:         "yes",
			b:         "no",
			expected:  "no",
		},
		{
			name:      "true condition with integers",
			condition: true,
			a:         1,
			b:         2,
			expected:  1,
		},
		{
			name:      "false condition with integers",
			condition: false,
			a:         1,
			b:         2,
			expected:  2,
		},
		{
			name:      "true condition with booleans",
			condition: true,
			a:         true,
			b:         false,
			expected:  true,
		},
		{
			name:      "false condition with booleans",
			condition: false,
			a:         true,
			b:         false,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch a := tt.a.(type) {
			case string:
				b := tt.b.(string)
				expected := tt.expected.(string)
				if got := If(tt.condition, a, b); got != expected {
					t.Errorf("If() = %v, want %v", got, expected)
				}
			case int:
				b := tt.b.(int)
				expected := tt.expected.(int)
				if got := If(tt.condition, a, b); got != expected {
					t.Errorf("If() = %v, want %v", got, expected)
				}
			case bool:
				b := tt.b.(bool)
				expected := tt.expected.(bool)
				if got := If(tt.condition, a, b); got != expected {
					t.Errorf("If() = %v, want %v", got, expected)
				}
			}
		})
	}
}
