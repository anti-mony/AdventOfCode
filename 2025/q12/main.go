package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent.of.code/util"
)

type shape [][]int
type region struct {
	rows, cols int
	qShapes    []int
}

type Input struct {
	shapes  []shape
	regions []region
}

type InputType Input

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	// for _, s := range inp.shapes {
	// 	util.PrintMatrix(s)
	// }

	// for _, r := range inp.regions {
	// 	fmt.Println(r)
	// }

	fmt.Println("Q1: ", Q1(inp))
	fmt.Println("Q2: ", Q2(inp))
}

func Q1(inp InputType) int {
	return -1
}

func Q2(inp InputType) int {
	return -1
}

func parseInput(filename string) (InputType, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return InputType{}, err
	}

	lineIdx := 0
	shapes := make([]shape, 0)
	regions := make([]region, 0)

	for lineIdx = 0; lineIdx < len(lines); lineIdx += 5 {
		if strings.Contains(lines[lineIdx], "x") {
			break
		}

		s := make([][]int, 3)
		for i := range 3 {
			s[i] = make([]int, 3)
		}

		si := 0
		for _, l := range lines[lineIdx+1 : lineIdx+4] {
			for li, v := range l {
				s[si][li] = 0
				if v == '#' {
					s[si][li] = 1
				}
			}
			si++
		}

		shapes = append(shapes, s)
	}

	for l := lineIdx; l < len(lines); l++ {
		splits := strings.Split(lines[l], ":")
		qS, err := util.SpaceSeparatedStringOfNumbersToIntSlice(splits[1])
		if err != nil {
			return InputType{}, err
		}
		splits2 := strings.Split(splits[0], "x")
		r := region{
			rows:    util.StringToNumber(splits2[0]),
			cols:    util.StringToNumber(splits2[1]),
			qShapes: qS,
		}

		regions = append(regions, r)
	}

	return InputType{
		shapes:  shapes,
		regions: regions,
	}, nil
}
