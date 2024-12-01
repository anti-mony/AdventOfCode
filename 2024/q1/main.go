package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	listA, listB, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer 1", solveQ1(listA, listB))
	fmt.Println("Answer 2", solveQ2(listA, listB))
}

func solveQ1(a, b []int) int {
	result := 0
	slices.Sort(a)
	slices.Sort(b)

	for i := range len(a) {
		result += util.Abs(a[i] - b[i])
	}

	return result
}

func solveQ2(a, b []int) int {
	result := 0

	freqA := list.Frequency(a)
	freqB := list.Frequency(b)

	for k, v := range freqA {
		if count, found := freqB[k]; found {
			result += k * count * v
		}
	}

	return result
}

func parseInput(filename string) ([]int, []int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}
	listA := make([]int, len(lines))
	listB := make([]int, len(lines))

	for i, line := range lines {
		nums, err := util.DelimitedStringOfNumbersToIntSlice(line)
		if err != nil {
			return nil, nil, err
		}
		listA[i] = nums[0]
		listB[i] = nums[1]
	}

	return listA, listB, nil
}
