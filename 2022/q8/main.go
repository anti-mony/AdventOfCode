package main

import (
	"fmt"
	"log"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	trees := solveP1(input)

	fmt.Println("Answer P1: ", len(trees)+2*(len(input)+len(input[0]))-4)

	fmt.Println("Answer P2: ", solveP2(input))
}

func solveP2(input [][]int) int {
	result := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			ss := getScenicScore(input, grid.NewCoordinate(i, j))
			if ss > result {
				// fmt.Printf("(%d, %d) -> %d \n", i, j, ss)
				result = ss
			}
		}
	}

	return result
}

func getScenicScore(input [][]int, start grid.Coordinate) int {
	scenicScore := 1

	for d := 1; d < 5; d++ {
		direction := grid.Direction(d)
		next := start.MoveTowards(direction)

		canSee := 0
		for inBound(input, next) {
			if input[start.X][start.Y] > input[next.X][next.Y] {
				canSee++
				next = next.MoveTowards(direction)
			} else {
				canSee++
				break
			}
		}

		scenicScore *= canSee
	}

	return scenicScore
}

func inBound(in [][]int, c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(in) && c.Y >= 0 && c.Y < len(in[0])
}

func solveP1(in [][]int) map[grid.Coordinate]bool {
	seen := make(map[grid.Coordinate]bool)

	// Ignore the last edge layer of the grid
	ROWS := len(in) - 1
	COLS := len(in[0]) - 1

	// Travel EAST --> WEST
	for i := 1; i < ROWS; i++ {
		maxSofar := in[i][0]
		for j := 1; j < COLS; j++ {
			if in[i][j] > maxSofar {
				seen[grid.NewCoordinate(i, j)] = true
				maxSofar = in[i][j]
			}
		}
	}

	// Travel WEST --> EAST
	for i := 1; i < ROWS; i++ {
		maxSoFar := in[i][COLS]
		for j := COLS - 1; j > 0; j-- {
			if in[i][j] > maxSoFar {
				seen[grid.NewCoordinate(i, j)] = true
				maxSoFar = in[i][j]
			}
		}
	}

	// Travel NORTH --> SOUTH
	for c := 1; c < COLS; c++ {
		maxSofar := in[0][c]
		for r := 1; r < ROWS; r++ {
			if in[r][c] > maxSofar {
				seen[grid.NewCoordinate(r, c)] = true
				maxSofar = in[r][c]
			}
		}
	}

	// Travel NORTH --> SOUTH
	for c := 1; c < COLS; c++ {
		maxSoFar := in[ROWS][c]
		for r := ROWS - 1; r > 0; r-- {
			if in[r][c] > maxSoFar {
				seen[grid.NewCoordinate(r, c)] = true
				maxSoFar = in[r][c]
			}
		}
	}

	return seen
}

func parseInput(filename string) ([][]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(lines))

	for i, line := range lines {
		result[i] = make([]int, len(line))
		for j, ch := range line {
			result[i][j] = util.StringToNumber(string(ch))
		}
	}

	return result, nil
}
