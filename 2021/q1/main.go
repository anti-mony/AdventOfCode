package main

import (
	"fmt"
	"log"

	"advent.of.code/util"
)

func main() {

	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("Answer Q1:%d\n", increases(input))
	fmt.Printf("Answer Q2:%d", increases(movingSum(input, 3)))
}

func movingSum(input []int, windowSize int) []int {
	result := make([]int, 0)
	i := 0
	for i < len(input)-windowSize+1 {
		s := 0
		for j := 0; j < windowSize; j++ {
			s += input[i+j]
		}
		result = append(result, s)
		i++
	}
	return result
}

func increases(input []int) int {
	result := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			result++
		}
	}

	return result
}

func parseInput(filename string) ([]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	result := make([]int, len(lines))
	for i, line := range lines {
		result[i] = util.StringToNumber(line)
	}
	return result, nil
}
