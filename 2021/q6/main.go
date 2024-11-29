package main

import (
	"fmt"
	"log"

	"advent.of.code/util"
)

const RESET = 6
const START_DAYS = 8

func main() {
	inp, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer P1", solveQ1(inp, 80))

	inp, err = parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	q2Ans := solveQ2(inp, 256)
	fmt.Println("Answer P2", q2Ans)
}

func solveQ1(starts []int, days int) int {
	for day := 0; day < days; day++ {
		nFishes := len(starts)
		for i := 0; i < nFishes; i++ {
			if starts[i] > 0 {
				starts[i] -= 1
			} else {
				starts = append(starts, START_DAYS)
				starts[i] = RESET
			}
		}
	}

	return len(starts)
}

func solveQ2(starts []int, days int) int {
	init := make(map[int]int)

	for _, s := range starts {
		if v, found := init[s]; found {
			init[s] = v + 1
		} else {
			init[s] = 1
		}
	}

	for day := 0; day < days; day++ {
		ninit := make(map[int]int)
		for timer, numberOfFishes := range init {
			if timer == 0 {
				ninit[RESET] += numberOfFishes
				ninit[START_DAYS] += numberOfFishes
			} else {
				ninit[timer-1] += numberOfFishes
			}
		}
		init = ninit
	}

	r := 0
	for _, v := range init {
		r += v
	}

	return r
}

func parseInput(filename string) ([]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	starts, err := util.DelimitedStringOfNumbersToIntSlice(lines[0])
	if err != nil {
		return nil, err
	}

	return starts, nil
}
