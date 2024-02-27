package main

import (
	"fmt"
	"log"
	"regexp"

	"advent.of.code/util"
)

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer P1: ", solveP1(input))
	fmt.Println("Answer P2: ", solveP2(input))
}

func solveP2(pairs []Pair) int {
	result := 0

	for _, pair := range pairs {

		smallerStart := pair.A
		biggerStart := pair.B
		if pair.A.Start > pair.B.Start {
			smallerStart = pair.B
			biggerStart = pair.A
		}

		if biggerStart.Start <= smallerStart.End {
			result++
		} else if biggerStart.End <= smallerStart.End {
			result++
		}

	}
	return result
}

func solveP1(pairs []Pair) int {
	result := 0

	for _, pair := range pairs {
		if pair.A.Start <= pair.B.Start && pair.A.End >= pair.B.End {
			result++
		} else if pair.B.Start <= pair.A.Start && pair.B.End >= pair.A.End {
			result++
		}
	}
	return result
}

type Section struct {
	Start int
	End   int
}

type Pair struct {
	A Section
	B Section
}

func parseInput(filename string) ([]Pair, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([]Pair, 0)

	re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		result = append(result, Pair{
			A: Section{util.StringToNumber(matches[1]), util.StringToNumber(matches[2])},
			B: Section{util.StringToNumber(matches[3]), util.StringToNumber(matches[4])},
		})
	}

	return result, nil
}
