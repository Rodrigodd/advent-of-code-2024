package day0

import (
	"testing"
)

type Results struct {
	part1 int
	part2 int
}

// Test functions (can coexist with the main program).
func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Results
	}{
		{
			name:     "Simple case",
			input:    ``,
			expected: Results{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1, part2 := solve(tt.input)
			result := Results{part1, part2}
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
