package main

import (
	"fmt"
	"log"
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
	fmt.Println("A2: ", Q2(maze))
}

func Q1(maze [][]string) int {
	start := grid.NewCoordinate(util.FindIndexMatrix(maze, "S"))
	end := grid.NewCoordinate(util.FindIndexMatrix(maze, "E"))
	walls := make([]grid.Coordinate, 0)

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == "#" {
				walls = append(walls, grid.NewCoordinate(i, j))
			}
		}
	}

	regularTime := FindShortestPath(maze, start, end)
	result := 0

	for _, wall := range walls {
		maze[wall.X][wall.Y] = "."
		r := FindShortestPath(maze, start, end)
		if regularTime-r >= 100 {
			result++
		}
		maze[wall.X][wall.Y] = "#"
	}

	return result
}

func Q2(maze [][]string) int {
	start := grid.NewCoordinate(util.FindIndexMatrix(maze, "S"))
	end := grid.NewCoordinate(util.FindIndexMatrix(maze, "E"))

	regularTime := FindShortestPath(maze, start, end)
	result := 0

	fmt.Println("regular time", regularTime)

	fmt.Println("cheat time", FindShortestPathWithCheats(maze, start, end))

	return result
}

type Key struct {
	c     grid.Coordinate
	steps int
	cheat *grid.Coordinate
}

func FindShortestPath(maze [][]string, start grid.Coordinate, end grid.Coordinate) int {
	seen := map[grid.Coordinate]int{}
	heap := list.NewHeap[Key](func(a, b Key) bool {
		return a.steps <= b.steps
	})

	heap.Push(Key{
		c:     start,
		steps: 0,
	})

	for heap.Len() > 0 {
		h := heap.Pop()
		if _, found := seen[h.c]; found {
			continue
		}
		if h.c == end {
			return h.steps
		}
		seen[h.c] = h.steps

		for _, delta := range grid.DIRECTIONS_STRAIGHT {
			n := h.c.Add(delta)
			_, found := seen[n]
			if inBound(maze, n) && !found && maze[n.X][n.Y] != "#" {
				heap.Push(Key{
					c:     n,
					steps: h.steps + 1,
				})
			}
		}
	}

	return 0
}

func FindShortestPathWithCheats(maze [][]string, start grid.Coordinate, end grid.Coordinate) int {
	seen := map[grid.Coordinate]int{}
	heap := list.NewHeap[Key](func(a, b Key) bool {
		return a.steps <= b.steps
	})

	heap.Push(Key{
		c:     start,
		steps: 0,
	})

	for heap.Len() > 0 {
		h := heap.Pop()
		if _, found := seen[h.c]; found {
			continue
		}

		seen[h.c] = h.steps

		for _, delta := range grid.DIRECTIONS_STRAIGHT {
			n := h.c.Add(delta)
			_, found := seen[n]
			if inBound(maze, n) && !found {
				if maze[n.X][n.Y] != "#" {
					heap.Push(Key{
						c:     n,
						steps: h.steps + 1,
					})
				} else {
					if h.cheat == nil {
						heap.Push(Key{
							c:     n,
							cheat: &n,
							steps: h.steps + 1,
						})
					} else {
						if h.cheat.RectilinearDistanceFrom(n) < 20 {
							heap.Push(Key{
								c:     n,
								cheat: h.cheat,
								steps: h.steps + 1})
						}
					}
				}
			}
		}
	}

	for k, v := range seen {
		maze[k.X][k.Y] = fmt.Sprintf("%d", v)
	}

	util.PrintMatrix(maze)

	return seen[end]
}

func inBound(maze [][]string, c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(maze) && c.Y >= 0 && c.Y < len(maze[c.X])
}
