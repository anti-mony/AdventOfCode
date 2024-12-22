package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	maze, err := util.ReadStringMatrixFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("A1: ", Q1(maze))
}

type key struct {
	c grid.Coordinate
	d grid.Direction
	n int
}

func Q1(maze [][]string) int {
	si, sj := util.FindIndexMatrix(maze, "S")
	ei, ej := util.FindIndexMatrix(maze, "E")
	return FindShortestPath(maze, grid.NewCoordinate(si, sj), grid.NewCoordinate(ei, ej))
}

func FindShortestPath(maze [][]string, start, end grid.Coordinate) int {
	seen := map[key]int{}

	hp := list.NewHeap[key](func(a, b key) bool {
		return a.n <= b.n
	})

	hp.Push(key{
		c: start,
		d: grid.DirectionEast,
		n: 0,
	})

	for hp.Len() > 0 {
		curr := hp.Pop()
		seen[key{
			c: curr.c,
			d: curr.d,
		}] = curr.n

		for dir := range grid.DIRECTIONS_STRAIGHT {
			if dir == curr.d.Reverse() {
				continue
			} else if dir != curr.d {
				_, found := seen[key{c: curr.c, d: dir}]
				if !found {
					hp.Push(key{
						c: curr.c,
						d: dir,
						n: curr.n + 1000,
					})
				}
			} else {
				next := curr.c.MoveTowards(curr.d)
				_, found := seen[key{c: next, d: dir}]
				if inBound(maze, next.X, next.Y) && !found && maze[next.X][next.Y] != "#" {
					hp.Push(key{
						c: next,
						d: dir,
						n: curr.n + 1,
					})
				}
			}
		}
	}

	minS := math.MaxInt
	for d := range grid.DIRECTIONS_STRAIGHT {
		minS = min(minS, seen[key{c: end, d: d}])
	}

	return minS
}

func inBound(maze [][]string, i, j int) bool {
	return i >= 0 && i < len(maze) && j >= 0 && j < len(maze[i])
}
