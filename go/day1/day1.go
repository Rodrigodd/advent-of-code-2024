package day1

import (
	"fmt"
	"github.com/Rodrigodd/advent-of-code-2024/util"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(input string) (int, int) {
	var lefts, rights []int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		splitted := strings.SplitN(line, " ", 2)

		ls := strings.TrimSpace(splitted[0])
		l, err := strconv.Atoi(ls)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %s to int: %v", splitted[0], err))
		}
		lefts = append(lefts, l)

		rs := strings.TrimSpace(splitted[1])
		r, err := strconv.Atoi(rs)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %s to int: %v", splitted[1], err))
		}
		rights = append(rights, r)
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	distance := 0

	for idx, l := range lefts {
		r := rights[idx]
		dist := util.AbsInt(l - r)
		distance += dist
	}

	fmt.Println("Part 1:", distance)

	// Part 2

	similarity := 0
	for _, l := range lefts {
		count := 0
		for _, r := range rights {
			if l == r {
				count += 1
			}
		}
		similarity += count * l
	}

	fmt.Println("Part 2:", similarity)

	return distance, similarity
}

func Day1() {
	input, err := os.ReadFile("../inputs/day1.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file: %v", err))
	}

	solve(string(input))
}
