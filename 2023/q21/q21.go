package main

import (
	"fmt"
	"log"

	"advent.of.code/grid"
	"advent.of.code/util"
)

const DEBUG = false

func main() {
	numberOfSteps := 6
	maze, startPos, err := getInput("input.small.txt")
	if err != nil {
		log.Fatal(err)
	}

	if DEBUG {
		length, width := len(maze), 0
		if length > 0 {
			width = len(maze[0])
		}
		grid.PrintGrid(maze)
		fmt.Printf("Number of Steps: %d | Maze Size: %d X %d | Start Position: %v\n", numberOfSteps, length, width, startPos)
	}

	fmt.Printf("P1 Answer is %d \n", solveP1(maze, startPos, numberOfSteps))
}

func solveP1(maze [][]string, start grid.Coordinate, numberOfSteps int) int {
	result := 0

	return result
}

func getInput(filename string) ([][]string, grid.Coordinate, error) {
	start := grid.Coordinate{}

	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, start, err
	}

	result := make([][]string, len(lines))

	for i, line := range lines {
		r := make([]string, len(line))
		for ii, c := range line {
			r[ii] = string(c)
			if string(c) == "S" {
				start = grid.Coordinate{X: i, Y: ii}
			}
		}
		result[i] = r
	}

	return result, start, nil
}
