package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Rodrigodd/advent-of-code-2024/util"
)

func isSafe(levels []int) bool {
	increasing := levels[0] < levels[1]
	for i := 1; i < len(levels); i++ {
		if increasing && levels[i-1] >= levels[i] {
			return false
		} else if !increasing && levels[i-1] <= levels[i] {
			return false
		} else {
			diff := util.AbsInt(levels[i-1] - levels[i])
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}

func solve(input string) (int, int) {
	safeCount := 0
	for _, report := range strings.Split(strings.TrimSpace(input), "\n") {
		levelsStr := strings.Split(report, " ")
		levels := make([]int, len(levelsStr))
		for i, levelStr := range levelsStr {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert %s to int: %v", levelStr, err))
			}
			levels[i] = level
		}

		if isSafe(levels) {
			safeCount += 1
		}
	}

	fmt.Println("Part 1:", safeCount)

	safeCount2 := 0
	for _, report := range strings.Split(strings.TrimSpace(input), "\n") {
		levelsStr := strings.Split(report, " ")
		levels := make([]int, len(levelsStr))
		for i, levelStr := range levelsStr {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert %s to int: %v", levelStr, err))
			}
			levels[i] = level
		}

		if isSafe(levels) {
			safeCount2 += 1
		} else {
			for i := 0; i < len(levels); i++ {
				dampened := append([]int{}, levels[:i]...)
				dampened = append(dampened, levels[i+1:]...)
				if isSafe(dampened) {
					safeCount2 += 1
					break
				}
			}
		}
	}

	fmt.Println("Part 2:", safeCount2)

	return safeCount, safeCount2
}

func Day2() {
	input, err := os.ReadFile("../inputs/day2.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file: %v", err))
	}

	solve(string(input))
}
