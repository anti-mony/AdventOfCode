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

const (
	HEIGHT = 70
	WIDTH  = 70
	NBYTES = 3011
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	// util.PrintMatrix(inp)

	fmt.Println("A1: ", FindMinSteps(inp))
}

type key struct {
	c     grid.Coordinate
	steps int
}

func FindMinSteps(memory [][]string) int {
	seen := map[grid.Coordinate]int{}
	start, end := grid.NewCoordinate(0, 0), grid.NewCoordinate(HEIGHT, WIDTH)

	heap := list.NewHeap[key](func(a, b key) bool {
		return a.steps <= b.steps
	})

	heap.Push(key{
		c:     start,
		steps: 0,
	})

	for heap.Len() > 0 {
		current := heap.Pop()
		if _, found := seen[current.c]; found {
			continue
		}

		seen[current.c] = current.steps

		for _, delta := range grid.DIRECTIONS_STRAIGHT {
			n := current.c.Add(delta)
			_, found := seen[n]
			if inBound(memory, n) && !found && memory[n.X][n.Y] != "#" {
				heap.Push(key{
					c:     n,
					steps: current.steps + 1,
				})
			}
		}
	}

	for k, v := range seen {
		memory[k.X][k.Y] = fmt.Sprintf("%d", v)
	}

	// util.PrintMatrix(memory)

	return seen[end]
}

func inBound(memory [][]string, c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(memory) && c.Y >= 0 && c.Y < len(memory[c.X])
}

func parseInput(filename string) ([][]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([][]string, HEIGHT+1)
	for i := range HEIGHT + 1 {
		result[i] = make([]string, WIDTH+1)
		for j := 0; j < WIDTH+1; j++ {
			result[i][j] = "."
		}
	}

	for _, line := range lines[:NBYTES] {
		splits := strings.Split(line, ",")
		x := util.StringToNumber(splits[0])
		y := util.StringToNumber(splits[1])
		result[y][x] = "#"
	}

	return result, nil
}
