package utils

import (
	"testing"
)

type TestJSONStruct struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Tags    []string `json:"tags"`
	Details struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"details"`
}

func TestJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name: "struct with nested fields",
			input: TestJSONStruct{
				Name: "John",
				Age:  30,
				Tags: []string{"developer", "golang"},
				Details: struct {
					City    string `json:"city"`
					Country string `json:"country"`
				}{
					City:    "New York",
					Country: "USA",
				},
			},
			expected: `{
    "name": "John",
    "age": 30,
    "tags": [
        "developer",
        "golang"
    ],
    "details": {
        "city": "New York",
        "country": "USA"
    }
}`,
		},
		{
			name:  "simple map",
			input: map[string]interface{}{"key": "value"},
			expected: `{
    "key": "value"
}`,
		},
		{
			name:  "array of strings",
			input: []string{"a", "b", "c"},
			expected: `[
    "a",
    "b",
    "c"
]`,
		},
		{
			name:     "nil value",
			input:    nil,
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JSON(tt.input)
			if got != tt.expected {
				t.Errorf("JSON() = %v, want %v", got, tt.expected)
			}
		})
	}
}
