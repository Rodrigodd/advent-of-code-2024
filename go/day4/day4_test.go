package day4

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
			name: "Simple case",
			input: `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			expected: Results{18, 9},
		},
		{
			name: "Simple case, part 2",
			input: `
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`,
			expected: Results{0, 9},
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
