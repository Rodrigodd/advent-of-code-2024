package day0

import (
	"fmt"
	"os"
)

func solve(input string) (int, int) {
	fmt.Println(input)
	return 0, 0
}

func Day0() {
	input, err := os.ReadFile("../inputs/day0.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file: %v", err))
	}

	solve(string(input))
}
