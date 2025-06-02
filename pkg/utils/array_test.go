package utils

import (
	"testing"
)

func TestContainsInArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      interface{}
		value    interface{}
		expected bool
	}{
		{
			name:     "string array contains value",
			arr:      []string{"a", "b", "c"},
			value:    "b",
			expected: true,
		},
		{
			name:     "string array does not contain value",
			arr:      []string{"a", "b", "c"},
			value:    "d",
			expected: false,
		},
		{
			name:     "int array contains value",
			arr:      []int{1, 2, 3},
			value:    2,
			expected: true,
		},
		{
			name:     "int array does not contain value",
			arr:      []int{1, 2, 3},
			value:    4,
			expected: false,
		},
		{
			name:     "empty array",
			arr:      []string{},
			value:    "a",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch arr := tt.arr.(type) {
			case []string:
				value := tt.value.(string)
				if got := ContainsInArray(arr, value); got != tt.expected {
					t.Errorf("ContainsInArray() = %v, want %v", got, tt.expected)
				}
			case []int:
				value := tt.value.(int)
				if got := ContainsInArray(arr, value); got != tt.expected {
					t.Errorf("ContainsInArray() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestMapToArray(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name: "string map to array",
			input: map[string]string{
				"a": "1",
				"b": "2",
				"c": "3",
			},
			expected: []string{"1", "2", "3"},
		},
		{
			name: "int map to array",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch m := tt.input.(type) {
			case map[string]string:
				expected := tt.expected.([]string)
				result := MapToArray(m)
				if len(result) != len(expected) {
					t.Errorf("MapToArray() length = %v, want %v", len(result), len(expected))
				}
				for _, v := range expected {
					if !ContainsInArray(result, v) {
						t.Errorf("MapToArray() missing value %v", v)
					}
				}
			case map[string]int:
				expected := tt.expected.([]int)
				result := MapToArray(m)
				if len(result) != len(expected) {
					t.Errorf("MapToArray() length = %v, want %v", len(result), len(expected))
				}
				for _, v := range expected {
					if !ContainsInArray(result, v) {
						t.Errorf("MapToArray() missing value %v", v)
					}
				}
			}
		})
	}
}
