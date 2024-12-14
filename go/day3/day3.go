package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solve(input string) (int, int) {
	mulOp := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	sum := 0
	sum2 := 0
	enabled := true
	for _, match := range mulOp.FindAllStringSubmatch(input, -1) {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += a * b
			if enabled {
				sum2 += a * b
			}
		}

	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)

	return sum, sum2
}

func Day3() {
	input, err := os.ReadFile("../inputs/day3.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file: %v", err))
	}

	solve(string(input))
}
