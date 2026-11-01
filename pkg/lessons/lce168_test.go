package lessons_test

import (
	"go-learnings/pkg/lessons"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		name         string
		columnNumber int
		expected     string
	}{
		{
			name:         "Example 4",
			columnNumber: 2147483647,
			expected:     "FXSHRXW",
		},
		{
			name:         "Example 1",
			columnNumber: 18278,
			expected:     "ZZZ",
		},
		{
			name:         "Example 1",
			columnNumber: 1,
			expected:     "A",
		},
		{
			name:         "Example 11",
			columnNumber: 26,
			expected:     "Z",
		},
		{
			name:         "Example 2",
			columnNumber: 28,
			expected:     "AB",
		},
		{
			name:         "Example 3",
			columnNumber: 701,
			expected:     "ZY",
		},
		{
			name:         "Example 31",
			columnNumber: 702,
			expected:     "ZZ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lessons.ConvertToTitle(tt.columnNumber)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
