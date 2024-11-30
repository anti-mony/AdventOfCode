package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
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

	res1, low := solveQ1(inp)

	fmt.Println("Answer Q1: ", res1)

	fmt.Println("Answer Q1: ", solveQ2(inp, low))

}

func solveQ1(inp *grid.Grid[int]) (int, []grid.Coordinate) {
	res := 0
	low := make([]grid.Coordinate, 0)

	R, C := inp.Dimensions()

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			c := grid.NewCoordinate(i, j)
			isSurrounded := true
			for _, delta := range grid.DIRECTIONS_STRAIGHT {
				n := c.Add(delta)
				if inp.InBound(n) {
					if inp.ValueAt(n) <= inp.ValueAt(c) {
						isSurrounded = false
					}
				}
			}
			if isSurrounded {
				res += inp.ValueAt(c) + 1
				low = append(low, c)
			}
		}
	}

	return res, low
}

func solveQ2(inp *grid.Grid[int], lows []grid.Coordinate) int {
	sizes := make([]int, len(lows))

	for i, low := range lows {
		size := 1
		q := list.NewQueue[grid.Coordinate]()
		q.Push(low)
		seen := map[grid.Coordinate]bool{
			low: true,
		}

		for q.Len() > 0 {
			c := q.Pop()
			for _, delta := range grid.DIRECTIONS_STRAIGHT {
				n := c.Add(delta)
				_, seenAlready := seen[n]
				if !seenAlready && inp.InBound(n) {
					if inp.ValueAt(c) < inp.ValueAt(n) && inp.ValueAt(n) < 9 {
						seen[n] = true
						q.Push(n)
						size++
					}
				}
			}
		}

		sizes[i] = size
	}

	slices.Sort(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func parseInput(filename string) (*grid.Grid[int], error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	return grid.NewIntGridFromStringSlice(lines)
}
