package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/util"
)

const (
	_paperRoll = "@"
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
	input, err := grid.NewStringGridFromMatrix[string](inp)
	if err != nil {
		panic(err)
	}

	result := 0

	dirs := grid.Directions()
	rows, cols := input.Dimensions()
	for i := range rows {
		for j := range cols {
			c := grid.NewCoordinate(i, j)
			if input.ValueAt(c) != _paperRoll {
				continue
			}
			rollCount := 0
			for _, d := range dirs {
				n := c.MoveTowards(d)
				if input.InBound(n) {
					if input.ValueAt(n) == _paperRoll {
						rollCount++
					}
				}
			}
			if rollCount < 4 {
				result++
			}
		}
	}

	return result
}

func Q2(inp InputType) int {
	return -1
}

func parseInput(filename string) (InputType, error) {
	return util.ReadStringMatrixFromFile(filename)
}
