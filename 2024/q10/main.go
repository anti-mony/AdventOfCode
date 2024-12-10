package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/util"
)

const (
	PEAK = 9
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := util.ReadIntMatrixFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	starts := []grid.Coordinate{}
	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp[i]); j++ {
			if inp[i][j] == 0 {
				starts = append(starts, grid.NewCoordinate(i, j))
			}
		}
	}

	q1, q2 := Q(inp, starts)
	fmt.Println("--1-->", q1)
	fmt.Println("--2-->", q2)
}

func Q(mountain [][]int, starts []grid.Coordinate) (int, int) {
	result1 := 0
	result2 := 0

	for _, start := range starts {
		seen := map[grid.Coordinate]bool{}
		result2 += reachPeak(mountain, start, seen)
		result1 += len(seen)
	}

	return result1, result2
}

func reachPeak(mountain [][]int, start grid.Coordinate, seen map[grid.Coordinate]bool) int {
	if mountain[start.X][start.Y] == PEAK {
		seen[start] = true
		return 1
	}

	result := 0
	for _, d := range grid.DIRECTIONS_STRAIGHT {
		nx, ny := start.X+d.X, start.Y+d.Y
		if inBound(mountain, nx, ny) && mountain[nx][ny] == mountain[start.X][start.Y]+1 {
			result += reachPeak(mountain, grid.NewCoordinate(nx, ny), seen)
		}
	}

	return result
}

func inBound(total [][]int, i, j int) bool {
	return i >= 0 && i < len(total) && j >= 0 && j < len(total[i])
}
