package main

import (
	"fmt"
	"log"
	"sort"

	"advent.of.code/util"
)

func main() {
	inp, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer P1: ", solveP1(inp))

	fmt.Println("Answer P2: ", solveP2(inp))

}

func solveP1(elves []Elf) int {
	maxCal := 0

	for _, elf := range elves {
		c := elf.GetTotalCalories()
		if c > maxCal {
			maxCal = c
		}
	}

	return maxCal
}

func solveP2(elves []Elf) int {

	sort.Slice(elves, func(i, j int) bool { return elves[i].GetTotalCalories() < elves[j].GetTotalCalories() })
	res := 0
	for i := len(elves) - 1; i > len(elves)-4; i-- {
		res += elves[i].GetTotalCalories()
	}

	return res

}

type Elf struct {
	Calories []int
}

func (e Elf) GetTotalCalories() int {
	sum := 0

	for _, c := range e.Calories {
		sum += c
	}

	return sum
}

func parseInput(filename string) ([]Elf, error) {
	inp := make([]Elf, 0)

	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	elf := Elf{Calories: make([]int, 0)}
	for _, line := range lines {
		elf.Calories = append(elf.Calories, util.StringToNumber(line))

		if line == "" {
			inp = append(inp, elf)
			elf = Elf{Calories: make([]int, 0)}
		}
	}

	return inp, nil
}
