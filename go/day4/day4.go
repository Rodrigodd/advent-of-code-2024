package day4

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Vec2 struct {
	x int
	y int
}

func solve(input string) (int, int) {
	puzzle := bytes.Split([]byte(strings.TrimSpace(input)), []byte("\n"))

	directions := []Vec2{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	fmt.Println("puzzle size:", len(puzzle), len(puzzle[0]))

	matchCount := 0

	for y := range len(puzzle) {
		for x := range len(puzzle[y]) {
			for _, d := range directions {
				cx := x
				cy := y
				match := true
			find:
				for _, r := range "XMAS" {
					c := byte(r)
					if cx < 0 || cx >= len(puzzle) || cy < 0 || cy >= len(puzzle[0]) {
						match = false
						break find
					}
					if puzzle[cx][cy] != c {
						match = false
						break find
					}
					cx += d.x
					cy += d.y
				}
				if match {
					matchCount += 1
				}
			}
		}
	}

	fmt.Println("Part1:", matchCount)

	patterns := [][]Vec2{
		{{1, 1}, {1, -1}},
		{{-1, -1}, {1, -1}},
		{{1, 1}, {-1, 1}},
		{{-1, -1}, {-1, 1}},
	}

	isMatch := func(x int, y int, dir1 Vec2, dir2 Vec2) bool {
		if puzzle[x][y] != 'A' {
			return false
		}
		if puzzle[x-dir1.x][y-dir1.y] != 'M' {
			return false
		}
		if puzzle[x+dir1.x][y+dir1.y] != 'S' {
			return false
		}
		if puzzle[x-dir2.x][y-dir2.y] != 'M' {
			return false
		}
		if puzzle[x+dir2.x][y+dir2.y] != 'S' {
			return false
		}
		return true
	}

	matchCount2 := 0

	for y := 1; y < len(puzzle)-1; y++ {
		for x := 1; x < len(puzzle[y])-1; x++ {
			for _, d := range patterns {
				if isMatch(x, y, d[0], d[1]) {
					matchCount2 += 1
				}
			}
		}
	}

	fmt.Println("Part2:", matchCount2)

	return matchCount, matchCount2
}

func Day4() {
	input, err := os.ReadFile("../inputs/day4.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file: %v", err))
	}

	solve(string(input))
}
