package day6

import (
	"fmt"
	"os"
	"strings"
)

type Tile int

const (
	Obstacle Tile = 1 << iota
	VisitedLeft
	VisitedRight
	VisitedUp
	VisitedDown
	NewObstacle
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var directionVectors = map[Direction]Vec2{
	Up:    {0, -1},
	Right: {1, 0},
	Down:  {0, 1},
	Left:  {-1, 0},
}

var visitedDirection = map[Direction]Tile{
	Up:    VisitedUp,
	Right: VisitedRight,
	Down:  VisitedDown,
	Left:  VisitedLeft,
}

func turnRight(dir Direction) Direction {
	return (dir + 1) % 4
}

const Visited = VisitedLeft | VisitedRight | VisitedUp | VisitedDown

type Vec2 struct {
	x, y int
}

// make a copy of the grid
func copyGrid(orGrid [][]Tile) [][]Tile {
	grid := make([][]Tile, len(orGrid))
	for y := range orGrid {
		grid[y] = make([]Tile, len(orGrid[y]))
		copy(grid[y], orGrid[y])
	}
	return grid
}

func willLoop(grid [][]Tile, startPos Vec2, startDir Direction) bool {
	guardPos := startPos
	guardDir := startDir
	//
	// printGrid(grid)

	isLoop := false
	for {
		guardVec := directionVectors[guardDir]
		nextPos := Vec2{guardPos.x + guardVec.x, guardPos.y + guardVec.y}
		if nextPos.y < 0 || nextPos.y >= len(grid) || nextPos.x < 0 || nextPos.x >= len(grid[nextPos.y]) {
			isLoop = false
			break
		}

		nextTile := grid[nextPos.y][nextPos.x]

		// move
		if nextTile&Obstacle != 0 {
			// rotate 90 degrees clockwise
			guardDir = turnRight(guardDir)
		} else {
			guardPos = nextPos
		}

		if grid[guardPos.y][guardPos.x]&visitedDirection[guardDir] != 0 {
			// same position and direction as before, will loop
			isLoop = true
			break
		}

		grid[guardPos.y][guardPos.x] |= visitedDirection[guardDir]
	}

	// if isLoop {
	// 	fmt.Println("Loop:", isLoop)
	// 	printGrid(grid)
	// }

	return isLoop
}

func solve(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]Tile, len(lines))
	guardPos := Vec2{-1, -1}
	guardDir := Up // assume it starts facing up
	for y := range grid {
		grid[y] = make([]Tile, len(lines[y]))
		for x, c := range lines[y] {
			switch c {
			case '^':
				guardPos = Vec2{x, y}
				grid[y][x] = VisitedUp
			case '#':
				grid[y][x] = Obstacle
			default:
				grid[y][x] = 0
			}
		}
	}

	startGrid := copyGrid(grid)
	startPos := guardPos
	startDir := guardDir

	// The tile the guard is on is already visited
	visitCount := 1
	timeTravelObstacleCount := 0

	for {
		// fmt.Println(guardPos, guardDir)
		// printGrid(grid)

		guardVec := directionVectors[guardDir]
		nextPos := Vec2{guardPos.x + guardVec.x, guardPos.y + guardVec.y}
		if nextPos.y < 0 || nextPos.y >= len(grid) || nextPos.x < 0 || nextPos.x >= len(grid[nextPos.y]) {
			break
		}

		nextTile := grid[nextPos.y][nextPos.x]

		if nextTile&Obstacle != 0 {
			// rotate 90 degrees clockwise
			guardDir = turnRight(guardDir)
		} else {
			// check if a obstacle here would cause a loop
			if grid[nextPos.y][nextPos.x]&NewObstacle == 0 {
				gridWithObstacle := copyGrid(startGrid)
				gridWithObstacle[nextPos.y][nextPos.x] = Obstacle | NewObstacle
				if willLoop(gridWithObstacle, startPos, startDir) {
					timeTravelObstacleCount++
					grid[nextPos.y][nextPos.x] |= NewObstacle
				}
			}

			guardPos = nextPos
			if nextTile&Visited == 0 {
				visitCount++
			}

			grid[guardPos.y][guardPos.x] |= visitedDirection[guardDir]
		}
	}
	//
	// printGrid(grid)

	fmt.Println("Part 1:", visitCount)
	fmt.Println("Part 2:", timeTravelObstacleCount)

	return visitCount, timeTravelObstacleCount
}

func Day6() {
	input, err := os.ReadFile("../inputs/day6.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file: %v", err))
	}

	solve(string(input))
}

func printGrid(grid [][]Tile) {
	fmt.Println("+" + strings.Repeat("-", len(grid[0])) + "+")
	for _, row := range grid {
		fmt.Print("|")
		for _, tile := range row {
			// switch tile {
			// case Visited:
			// 	fmt.Print("X")
			// case Obstacle:
			// 	fmt.Print("#")
			// default:
			// 	fmt.Print(".")
			// }
			switch tile {
			case 0:
				fmt.Print(".")
			case NewObstacle | Obstacle:
				fmt.Print("O")
			case Obstacle:
				fmt.Print("#")
			case VisitedUp:
				fmt.Print("^")
			case VisitedRight:
				fmt.Print(">")
			case VisitedDown:
				fmt.Print("v")
			case VisitedLeft:
				fmt.Print("<")
			case VisitedLeft | VisitedRight:
				fmt.Print("-")
			case VisitedUp | VisitedDown:
				fmt.Print("|")
			default:
				fmt.Print("+")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+" + strings.Repeat("-", len(grid[0])) + "+")
}
