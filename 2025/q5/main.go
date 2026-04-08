package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
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
	ranges := inp.FreshRanges
	slices.SortFunc(ranges, func(x, y []int) int { return cmp.Compare(x[0], y[0]) })
	current := ranges[0]
	merged := make([][]int, 0, len(ranges))
	for _, interval := range ranges[1:] {
		if current[1] < interval[0] {
			merged = append(merged, current)
			current = interval
		} else {
			current[1] = max(current[1], interval[1])
		}
	}
	merged = append(merged, current)

	result := 0
	for _, r := range merged {
		result += r[1] - r[0] + 1
	}

	return result
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
