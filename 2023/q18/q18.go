package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

type Task struct {
	Direction grid.Direction
	Steps     int
	Color     string
}

func main() {
	fileName := os.Args[1]
	g, err := getInput(fileName)
	if err != nil {
		log.Fatal(err, g)
	}

	fmt.Printf("Answer 1 solution is %d\n", solveP1(g))
}

func solveP1(input []Task) int {

	max := map[grid.Direction]int{
		grid.DirectionEast:  0,
		grid.DirectionWest:  0,
		grid.DirectionNorth: 0,
		grid.DirectionSouth: 0,
	}
	for _, t := range input {
		max[t.Direction] += t.Steps
	}

	maxRows := 2*(max[grid.DirectionNorth]+max[grid.DirectionSouth]) + 1
	maxCols := 2*(max[grid.DirectionEast]+max[grid.DirectionWest]) + 1

	ground := make([][]int, maxRows)
	for i := 0; i < maxRows; i++ {
		ground[i] = make([]int, maxCols)
	}

	pos := grid.Coordinate{maxRows/2 - 1, maxCols/2 - 1}
	ground[pos.X][pos.Y] = 1

	for _, task := range input {
		for i := 0; i < task.Steps; i++ {
			pos = pos.MoveTowards(task.Direction)
			ground[pos.X][pos.Y] = 1
		}
	}

	ground = pruneRows(ground)
	ground = pruneCols(ground)

	fmt.Printf("Grid Size: %dx%d \n", len(ground), len(ground[0]))

	return calculateVolume(ground, 1)
}

func pruneRows(in [][]int) [][]int {
	lastIndex := len(in) - 1
	for lastIndex >= 0 {
		for j := 0; j <= len(in)-1; j++ {
			if in[lastIndex][j] == 1 {
				lastIndex = -1
				break
			}
		}
		if lastIndex > 0 {
			in = in[:lastIndex]
			lastIndex = len(in) - 1
		}
	}

	firstIndex := 0
	for firstIndex < len(in) {
		for j := 0; j <= len(in[firstIndex])-1; j++ {
			if in[firstIndex][j] == 1 {
				return in
			}
		}
		in = in[firstIndex+1:]
		firstIndex = 0
	}

	return in
}

func pruneCols(in [][]int) [][]int {
	nRows := len(in)
	nCols := len(in[0])

	firstIndex := 0
	found := false
	for firstIndex < nCols {
		for j := 0; j < nRows; j++ {
			if in[j][firstIndex] == 1 {
				found = true
				break
			}
		}
		if found {
			break
		}
		firstIndex++
	}

	lastIndex := nCols - 1
	found = false
	for lastIndex > 0 {
		for j := 0; j < nRows; j++ {
			if in[j][lastIndex] == 1 {
				found = true
				break
			}
		}
		if found {
			break
		}
		lastIndex--
	}

	for i := 0; i < nRows; i++ {
		in[i] = in[i][firstIndex : lastIndex+1]
	}

	return in
}

func calculateVolume(ground [][]int, depth int) int {
	q := list.NewQueue()
	for i := 0; i < len(ground[0]); i++ {
		if ground[0][i] == 1 {
			q.Push(grid.Coordinate{1, i + 1})
			break
		}
	}

	for q.Len() > 0 {
		c := q.Pop()
		pos := c.(grid.Coordinate)
		for i := 1; i < 5; i++ {
			newPos := pos.MoveTowards(grid.Direction(i))
			if newPos.X >= 0 && newPos.X < len(ground) && newPos.Y >= 0 && newPos.Y < len(ground[newPos.X]) {
				if ground[newPos.X][newPos.Y] != 1 {
					ground[newPos.X][newPos.Y] = 1
					q.Push(newPos)
				}
			}
		}
	}

	grid.PrintGrid(ground)

	area := 0

	for i := 0; i < len(ground); i++ {
		for j := 0; j < len(ground[i]); j++ {
			if ground[i][j] == 1 {
				area++
			}
		}
	}

	return area

}

func moveToDirection(in string) grid.Direction {
	switch in {
	case "R":
		return grid.DirectionEast
	case "L":
		return grid.DirectionWest
	case "D":
		return grid.DirectionSouth
	}
	return grid.DirectionNorth
}

func getInput(filename string) ([]Task, error) {

	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([]Task, len(lines))

	for i, line := range lines {
		splits := strings.Split(line, " ")
		result[i] = Task{
			Direction: moveToDirection(splits[0]),
			Steps:     util.StringToNumber(splits[1]),
			Color:     splits[2],
		}
	}

	return result, nil
}
