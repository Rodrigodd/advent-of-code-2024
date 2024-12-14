package day2

import (
	"testing"
)

type Result struct {
	safeCount  int
	safeCount2 int
}

// Test functions (can coexist with the main program).
func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Result
	}{
		{
			name: "Simple case",
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			expected: Result{2, 4},
		},
		{
			name: "Other case",
			input: `
2 1 3 4
2 1 3 0
4 5 3 2
4 5 3 6
1 2 1
2 1 2
1 2 2
2 2 4`,
			expected: Result{0, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			safeCount, safeCount2 := solve(tt.input)
			if safeCount != tt.expected.safeCount || safeCount2 != tt.expected.safeCount2 {
				t.Errorf("Expected %d, got %d", tt.expected, Result{safeCount, safeCount2})
			}
		})
	}
}
