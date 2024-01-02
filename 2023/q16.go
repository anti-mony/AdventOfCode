package main

import (
	"fmt"
	"slices"
)

type VisitedQ16 struct {
	d map[Direction]bool
}

func q16sol() error {

	input, err := getInputQ16("input16.txt")
	if err != nil {
		return err
	}

	// print2DArray(input)

	fmt.Printf("Q16 P1 Answer is: %d \n", solveQ16P1(input, Coordinate{0, 0}, DirectionEast))

	fmt.Printf("Q16 P2 Answer is: %d \n", solveQ16P2(input))

	return nil
}

func solveQ16P1(in [][]string, start Coordinate, direction Direction) int {

	visited := make([][]VisitedQ16, len(in))
	for i := 0; i < len(in); i++ {
		r := make([]VisitedQ16, len(in[0]))
		for j := 0; j < len(in[0]); j++ {
			r[j].d = make(map[Direction]bool)
		}
		visited[i] = r
	}
	solveRecorsivelyQ16(in, visited, start, direction)

	result := 0

	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0]); j++ {
			if len(visited[i][j].d) > 0 {
				result += 1
			}
		}
	}

	return result
}

func solveQ16P2(in [][]string) int {

	energized := make([]int, 0)

	// Top row going south
	for i := 0; i < len(in[0]); i++ {
		energized = append(energized, solveQ16P1(in, Coordinate{0, i}, DirectionSouth))
	}

	// Bottom row going north
	for i := 0; i < len(in[0]); i++ {
		energized = append(energized, solveQ16P1(in, Coordinate{len(in) - 1, i}, DirectionNorth))
	}

	// First column going east
	for i := 0; i < len(in); i++ {
		energized = append(energized, solveQ16P1(in, Coordinate{i, 0}, DirectionEast))
	}

	// last column going west
	for i := 0; i < len(in); i++ {
		energized = append(energized, solveQ16P1(in, Coordinate{i, len(in[0]) - 1}, DirectionWest))
	}

	return slices.Max(energized)
}

func solveRecorsivelyQ16(in [][]string, visited [][]VisitedQ16, c Coordinate, direction Direction) {
	if _, ok := visited[c.x][c.y].d[direction]; ok {
		return
	}

	for {
		if !IsValidIndexQ16(in, c) {
			return
		}
		visited[c.x][c.y].d[direction] = true
		if (direction == DirectionEast || direction == DirectionWest) && in[c.x][c.y] == "|" {
			solveRecorsivelyQ16(in, visited, c, DirectionNorth)
			solveRecorsivelyQ16(in, visited, c, DirectionSouth)
			return
		}
		if (direction == DirectionNorth || direction == DirectionSouth) && in[c.x][c.y] == "-" {
			solveRecorsivelyQ16(in, visited, c, DirectionEast)
			solveRecorsivelyQ16(in, visited, c, DirectionWest)
			return
		}

		if in[c.x][c.y] == "/" {
			switch direction {
			case DirectionNorth:
				direction = DirectionEast
			case DirectionEast:
				direction = DirectionNorth
			case DirectionWest:
				direction = DirectionSouth
			case DirectionSouth:
				direction = DirectionWest
			}
		} else if in[c.x][c.y] == "\\" {
			switch direction {
			case DirectionNorth:
				direction = DirectionWest
			case DirectionEast:
				direction = DirectionSouth
			case DirectionWest:
				direction = DirectionNorth
			case DirectionSouth:
				direction = DirectionEast
			}
		}
		c = c.Add(DIRECTIONS[direction])
	}
}

func IsValidIndexQ16(in [][]string, c Coordinate) bool {
	return c.x >= 0 && c.x < len(in) && c.y >= 0 && c.y < len(in[0])
}

func getInputQ16(filaname string) ([][]string, error) {
	lines, err := getFileAsListOfStrings(filaname)
	if err != nil {
		return nil, err
	}

	result := make([][]string, len(lines))
	for i, l := range lines {
		row := make([]string, len(l))
		for j, c := range l {
			row[j] = string(c)
		}
		result[i] = row
	}
	return result, nil
}
