package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	left  int
	right int
}

func parseRules(input string) []Rule {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := make([]Rule, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "|")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		rules[i] = Rule{left, right}
	}

	return rules
}

func parseUpdates(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	updates := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		update := make([]int, len(parts))
		for j, part := range parts {
			fmt.Sscanf(part, "%d", &update[j])
		}
		updates[i] = update
	}
	return updates
}

// returns true if page1 < page2
func cmpByRule(page1, page2 int, rules []Rule) bool {
	for _, rule := range rules {
		if page1 == rule.left && page2 == rule.right {
			return true
		}
		if page1 == rule.right && page2 == rule.left {
			return false
		}
	}
	return false
}

func solve(input string) (int, int) {
	splitted := strings.SplitN(input, "\n\n", -1)
	rules := parseRules(splitted[0])
	updates := parseUpdates(splitted[1])

	part1 := 0
	part2 := 0

	for _, update := range updates {
		isValid := true
		for _, rule := range rules {
			leftIdx := -1
			rightIdx := -1
			for i, page := range update {
				if page == rule.left {
					leftIdx = i
				}
				if page == rule.right {
					rightIdx = i
				}
			}
			if leftIdx == -1 || rightIdx == -1 {
				continue
			}
			if leftIdx < rightIdx {
				continue
			}
			isValid = false
			break
		}
		if isValid {
			middlePage := len(update) / 2
			part1 += update[middlePage]
		} else {
			// bubble sort pages
			for i := 0; i < len(update); i++ {
				for j := 0; j < len(update)-i-1; j++ {
					if cmpByRule(update[j], update[j+1], rules) {
						update[j], update[j+1] = update[j+1], update[j]
					}
				}
			}
			middlePage := len(update) / 2
			part2 += update[middlePage]
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)

	return part1, part2
}

func Day5() {
	input, err := os.ReadFile("../inputs/day5.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file: %v", err))
	}

	solve(string(input))
}
