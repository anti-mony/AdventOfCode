package main

import (
	"fmt"
	"log"
	"math"

	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	inp, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solveQ1(inp))

	fmt.Println(solveQ2(inp))
}

func solveQ1(input []int) int {
	positions := map[int]int{}
	minP, maxP := math.MaxInt, math.MinInt
	for _, p := range input {
		if v, found := positions[p]; found {
			positions[p] = v + 1
		} else {
			positions[p] = 1
		}
		minP = min(minP, p)
		maxP = max(maxP, p)
	}

	costs := make([]int, 0)

	for i := minP; i < maxP; i++ {
		cost := 0
		for p, n := range positions {
			cost += util.Abs(p-i) * n
		}
		costs = append(costs, cost)
	}

	minVal, _ := list.Min(costs)

	return minVal + minP
}

func solveQ2(input []int) int {
	positions := map[int]int{}
	minP, maxP := math.MaxInt, math.MinInt
	for _, p := range input {
		if v, found := positions[p]; found {
			positions[p] = v + 1
		} else {
			positions[p] = 1
		}
		minP = min(minP, p)
		maxP = max(maxP, p)
	}

	costs := make([]int, 0)

	for i := minP; i < maxP; i++ {
		cost := 0
		for p, n := range positions {
			distance := util.Abs(p - i)
			cost += (distance) * (distance + 1) * n / 2
		}
		costs = append(costs, cost)
	}

	minVal, _ := list.Min(costs)

	return minVal + minP
}

func parseInput(filename string) ([]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	return util.DelimitedStringOfNumbersToIntSlice(lines[0])
}
