package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, x, y, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	q1a, seen := Q1(inp, x, y)
	fmt.Println(q1a)

	fmt.Println(Q2(inp, x, y, seen))
}

type key struct {
	C grid.Coordinate
	D grid.Direction
}

var dirmap = map[grid.Direction]grid.Direction{
	grid.DirectionNorth: grid.DirectionEast,
	grid.DirectionEast:  grid.DirectionSouth,
	grid.DirectionSouth: grid.DirectionWest,
	grid.DirectionWest:  grid.DirectionNorth,
}

func Q1(lab [][]string, x, y int) (int, map[grid.Coordinate]bool) {
	seen := map[key]bool{}
	seen2 := map[grid.Coordinate]bool{}
	direction := grid.DirectionNorth
	dx, dy := grid.DIRECTIONS[direction].X, grid.DIRECTIONS[direction].Y

	for {
		seen[key{grid.NewCoordinate(x, y), direction}] = true
		seen2[grid.NewCoordinate(x, y)] = true
		nx := x + dx
		ny := y + dy
		if !inbound(lab, nx, ny) {
			return len(seen2), seen2
		}
		if lab[nx][ny] == "#" {
			direction = dirmap[direction]
			dx, dy = grid.DIRECTIONS[direction].X, grid.DIRECTIONS[direction].Y
		} else {
			x, y = nx, ny
		}
	}

}

func Q2(lab [][]string, x, y int, already map[grid.Coordinate]bool) int {
	loops := 0
	delete(already, grid.NewCoordinate(x, y))
	for k := range already {
		lab[k.X][k.Y] = "#"
		if loop(lab, x, y) {
			loops++
		}
		lab[k.X][k.Y] = "."
	}

	return loops
}

func loop(lab [][]string, x, y int) bool {
	seen := map[key]bool{}
	direction := grid.DirectionNorth
	dx, dy := grid.DIRECTIONS[direction].X, grid.DIRECTIONS[direction].Y

	for {
		if _, found := seen[key{grid.NewCoordinate(x, y), direction}]; found {
			return true
		}
		seen[key{grid.NewCoordinate(x, y), direction}] = true
		nx := x + dx
		ny := y + dy
		if !inbound(lab, nx, ny) {
			return false
		}

		if lab[nx][ny] == "#" {
			direction = dirmap[direction]
			dx, dy = grid.DIRECTIONS[direction].X, grid.DIRECTIONS[direction].Y
		} else {
			x, y = nx, ny
		}
	}
}

func inbound(lab [][]string, x, y int) bool {
	return x >= 0 && x < len(lab) && y >= 0 && y < len(lab[x])
}

func parseInput(filename string) ([][]string, int, int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, -1, -1, err
	}
	sx, sy := -1, -1
	lab := make([][]string, len(lines))
	for i, line := range lines {
		lab[i] = util.StringToCharSlice(line)
		for j := 0; j < len(lab[i]); j++ {
			if lab[i][j] == "^" {
				sx, sy = i, j
			}
		}
	}

	return lab, sx, sy, nil
}
