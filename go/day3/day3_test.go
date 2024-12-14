package day3

import (
	"testing"
)

type Results struct {
	sum  int
	sum2 int
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
			input:    `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			expected: Results{161, 161},
		},
		{
			name:     "Simple case, part 2",
			input:    `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			expected: Results{161, 48},
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
