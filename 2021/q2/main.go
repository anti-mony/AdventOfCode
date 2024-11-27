package main

import (
	"fmt"
	"log"
	"strings"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {

	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	f := finalDestination(input)
	fmt.Printf("Answer Q1 (%v) : %d\n", f.String(), f.X*f.Y)

	f = finalDestination2(input)
	fmt.Printf("Answer Q2 (%v) : %d\n", f.String(), f.X*f.Y)
}

func finalDestination(input []DirVal) grid.Coordinate {
	init := grid.NewCoordinate(0, 0)
	for _, m := range input {
		init = init.MoveNUnitsTowards(m.Direction, m.Value)
	}
	return init
}

func finalDestination2(input []DirVal) grid.Coordinate {
	init := grid.NewCoordinate(0, 0)
	depth := 0
	for _, m := range input {
		init = init.MoveNUnitsTowards(m.Direction, m.Value)
		if m.Direction == grid.DirectionEast {
			depth += m.Value * init.X
		}
	}
	return grid.NewCoordinate(init.Y, depth)
}

type DirVal struct {
	Direction grid.Direction
	Value     int
}

func parseInput(filename string) ([]DirVal, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	result := make([]DirVal, len(lines))
	for i, line := range lines {
		splits := strings.Split(line, " ")
		switch splits[0] {
		case "forward":
			result[i] = DirVal{
				Direction: grid.DirectionEast,
				Value:     util.StringToNumber(splits[1]),
			}
		case "up":
			result[i] = DirVal{
				Direction: grid.DirectionNorth,
				Value:     util.StringToNumber(splits[1]),
			}
		case "down":
			result[i] = DirVal{
				Direction: grid.DirectionSouth,
				Value:     util.StringToNumber(splits[1]),
			}
		}
	}
	return result, nil
}
