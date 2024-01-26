package main

import (
	"fmt"
	"log"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {
	numberOfSteps := 64
	maze, startPos, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	length, width := len(maze), 0
	if length > 0 {
		width = len(maze[0])
	}
	fmt.Printf("Number of Steps: %d | Maze Size: %d X %d | Start Position: %v\n\n", numberOfSteps, length, width, startPos)

	fmt.Printf("P1 Answer is %d for %d steps \n", solveP1(maze, startPos, numberOfSteps), numberOfSteps)
}

func solveP1(maze [][]string, start grid.Coordinate, maxSteps int) int {

	visited := map[grid.Coordinate]bool{start: true}

	for i := 0; i < maxSteps; i++ {
		newVisited := map[grid.Coordinate]bool{start: true}

		for pos, _ := range visited {
			for i := 1; i <= 4; i++ {
				nextPos := pos.MoveTowards(grid.Direction(i))
				if _, alreadyVisited := visited[nextPos]; alreadyVisited ||
					!isValidCoordinate(maze, nextPos) ||
					maze[nextPos.X][nextPos.Y] == "#" {
					continue
				}
				newVisited[nextPos] = true
			}
		}
		visited = newVisited
	}

	return len(visited)
}

func isValidCoordinate(maze [][]string, c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(maze) && c.Y >= 0 && c.Y < len(maze[c.X])
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
