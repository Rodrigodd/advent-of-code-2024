package day1

import (
	"testing"
)

type Results struct {
	distance   int
	similarity int
}

// Test functions (can coexist with the main program).
func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Results
	}{
		{
			name: "Simple case",
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expected: Results{11, 31},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance, similarity := solve(tt.input)
			if distance != tt.expected.distance {
				t.Errorf("Expected %d, got %d", tt.expected, distance)
			}
			if similarity != tt.expected.similarity {
				t.Errorf("Expected %d, got %d", tt.expected, similarity)
			}
		})
	}
}
