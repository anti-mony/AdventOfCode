package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent.of.code/util"
)

type Input struct {
	FreshRanges [][]int
	Ingredients []int
}

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

func Q1(inp Input) int {
	result := 0

	for _, i := range inp.Ingredients {
		isFresh := false

		for _, rng := range inp.FreshRanges {
			if i >= rng[0] && i <= rng[1] {
				isFresh = true
				break
			}
		}

		if isFresh {
			result++
		}
	}

	return result
}

func Q2(inp Input) int {
	return -1
}

func parseInput(filename string) (Input, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return Input{}, err
	}

	ranges := make([][]int, 0, len(lines)/2)
	ingredients := make([]int, 0, len(lines)/2)

	i := 0
	line := "START"
	for i, line = range lines {
		if line == "" {
			break
		}
		r, err := util.SpaceSeparatedStringOfNumbersToIntSlice(strings.ReplaceAll(line, "-", " "))
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, r)
	}

	for _, line := range lines[i+1:] {
		ingredients = append(ingredients, util.StringToNumber(line))
	}

	return Input{
		FreshRanges: ranges,
		Ingredients: ingredients,
	}, nil
}
