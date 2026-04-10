package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

type InputType [][]string

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Q1: ", Q1(inp))
	fmt.Println("Q2: ", Q2(inp))
}

func Q1(inp InputType) int {
	manifold, _ := grid.NewStringGridFromMatrix(util.CopyMatrix(inp))
	splits := 0
	q := list.NewQueue[grid.Coordinate]()
	seen := list.NewSet[grid.Coordinate]()

	for i, v := range inp[0] {
		if v == "S" {
			c := grid.NewCoordinate(0, i)
			q.Push(c)
			seen.Add(c)
			break
		}
	}
	for q.Len() > 0 {
		l := q.Len()
		for range l {
			current := q.Pop()
			manifold.SetValueAt(current, "|")
			nxt := current.MoveTowards(grid.DirectionSouth)
			if !manifold.InBound(nxt) {
				continue
			}
			if manifold.ValueAt(nxt) == "^" {
				splits++
				l, r := grid.NewCoordinate(nxt.X, nxt.Y-1), grid.NewCoordinate(nxt.X, nxt.Y+1)
				if manifold.InBound(l) && !seen.Contains(l) {
					q.Push(l)
					seen.Add(l)
				}
				if manifold.InBound(r) && !seen.Contains(r) {
					q.Push(r)
					seen.Add(r)
				}
			} else if !seen.Contains(nxt) {
				q.Push(nxt)
				seen.Add(nxt)
			}
		}
	}

	return splits
}

func Q2(inp InputType) int {
	manifold, _ := grid.NewStringGridFromMatrix(inp)
	q := list.NewQueue[grid.Coordinate]()
	seen := make(map[grid.Coordinate]int)

	for i, v := range inp[0] {
		if v == "S" {
			c := grid.NewCoordinate(0, i)
			q.Push(c)
			seen[c] = 1
			break
		}
	}
	for q.Len() > 0 {
		l := q.Len()
		for range l {
			current := q.Pop()
			manifold.SetValueAt(current, "|")
			nxt := current.MoveTowards(grid.DirectionSouth)
			if !manifold.InBound(nxt) {
				continue
			}
			if manifold.ValueAt(nxt) == "^" {
				l, r := grid.NewCoordinate(nxt.X, nxt.Y-1), grid.NewCoordinate(nxt.X, nxt.Y+1)
				if manifold.InBound(l) {
					if _, ok := seen[l]; !ok {
						q.Push(l)
					}
					seen[l] += seen[current]
				}
				if manifold.InBound(r) {
					if _, ok := seen[r]; !ok {
						q.Push(r)
					}
					seen[r] += seen[current]
				}
			} else {
				if _, ok := seen[nxt]; !ok {
					q.Push(nxt)
				}
				seen[nxt] += seen[current]
			}
		}
	}

	r, c := manifold.Dimensions()
	timelines := 0

	for i := 0; i < c; i++ {
		timelines += seen[grid.NewCoordinate(r-1, i)]
	}

	return timelines

}

func parseInput(filename string) (InputType, error) {
	lines, err := util.ReadStringMatrixFromFile(filename)
	if err != nil {
		return nil, err
	}
	return lines, nil
}
