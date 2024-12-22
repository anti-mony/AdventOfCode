package main

import (
	"fmt"
	"log"
	"math"
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

	maze, err := util.ReadStringMatrixFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	steps, paths := Q1(maze)

	fmt.Println("A1: ", steps, paths)
	// for k, v := range adjList {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println("A2: ", Q2(maze, adjList))
}

type key struct {
	c grid.Coordinate
	d grid.Direction
}

type heapKey struct {
	c    grid.Coordinate
	d    grid.Direction
	n    int
	path []grid.Coordinate
}

func Q1(maze [][]string) (int, int) {
	s := grid.NewCoordinate(util.FindIndexMatrix(maze, "S"))
	e := grid.NewCoordinate(util.FindIndexMatrix(maze, "E"))
	return FindShortestPath(maze, s, e)
}

func Q2(maze [][]string, adjList map[grid.Coordinate][]grid.Coordinate) int {
	s := grid.NewCoordinate(util.FindIndexMatrix(maze, "S"))
	e := grid.NewCoordinate(util.FindIndexMatrix(maze, "E"))
	fmt.Println(e, s, len(adjList))

	result := 0
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if _, found := adjList[grid.NewCoordinate(i, j)]; found {
				result++
			}
		}
	}

	return result
}

func FindShortestPath(maze [][]string, start, end grid.Coordinate) (int, int) {
	seen := map[key]int{}
	path := map[grid.Coordinate]bool{}
	hp := list.NewHeap[heapKey](func(a, b heapKey) bool {
		return a.n <= b.n
	})

	hp.Push(heapKey{
		c:    start,
		d:    grid.DirectionEast,
		n:    0,
		path: []grid.Coordinate{start},
	})

	for hp.Len() > 0 {
		curr := hp.Pop()

		k := key{
			c: curr.c,
			d: curr.d,
		}

		if _, found := seen[key{}]; found {
			continue
		}

		if curr.c == end {
			for _, v := range curr.path {
				path[v] = true
			}
		}

		seen[k] = curr.n

		for dir := range grid.DIRECTIONS_STRAIGHT {
			if dir == curr.d.Reverse() {
				continue
			} else if dir != curr.d {
				_, found := seen[key{c: curr.c, d: dir}]
				if !found {
					hp.Push(heapKey{
						c:    curr.c,
						d:    dir,
						n:    curr.n + 1000,
						path: slices.Clone(curr.path),
					})
				}
			} else {
				next := curr.c.MoveTowards(curr.d)
				_, found := seen[key{c: next, d: dir}]
				if inBound(maze, next.X, next.Y) && !found && maze[next.X][next.Y] != "#" {
					hp.Push(heapKey{
						c:    next,
						d:    dir,
						n:    curr.n + 1,
						path: append(slices.Clone(curr.path), next),
					})
				}
			}
		}
	}

	minS := math.MaxInt
	for d := range grid.DIRECTIONS_STRAIGHT {
		minS = min(minS, seen[key{c: end, d: d}])
	}

	util.PrintMatrix(maze)
	for k := range path {
		maze[k.X][k.Y] = "O"
	}
	util.PrintMatrix(maze)

	return minS, len(path)
}

func inBound(maze [][]string, i, j int) bool {
	return i >= 0 && i < len(maze) && j >= 0 && j < len(maze[i])
}
