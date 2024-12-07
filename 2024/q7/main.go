package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"advent.of.code/util"
)

type Equation struct {
	Value   int
	Numbers []int
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

	fmt.Println(Q1(inp))

	fmt.Println(Q2(inp))
}

func Q1(equations []Equation) int {
	result := int(0)

	for _, equation := range equations {
		if isPossible(equation.Value, equation.Numbers[0], equation.Numbers[1:]) {
			result += equation.Value
		}
	}
	return result
}

func isPossible(answer int, current int, left []int) bool {
	if answer == int(current) {
		return true
	}
	if len(left) == 0 {
		return false
	}

	return isPossible(answer, current*left[0], left[1:]) || isPossible(answer, current+left[0], left[1:])
}

func Q2(equations []Equation) int {
	result := int(0)

	for _, equation := range equations {
		if isPossibleWithConcat(equation.Value, int(equation.Numbers[0]), equation.Numbers[1:], make([]string, 0)) {
			result += int(equation.Value)
		}
	}
	return result
}

func isPossibleWithConcat(answer int, current int, left []int, ops []string) bool {
	// fmt.Println(answer, current, left)
	if answer == current && len(left) == 0 {
		// For debuging P2
		// fmt.Println(ops, answer)
		return true
	}
	if len(left) == 0 {
		return false
	}

	return isPossibleWithConcat(answer, concat(current, left[0]), left[1:], append(ops, "||")) ||
		isPossibleWithConcat(answer, current*int(left[0]), left[1:], append(ops, "*")) ||
		isPossibleWithConcat(answer, current+int(left[0]), left[1:], append(ops, "+"))

}

func concat(a int, b int) int {
	num, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return num
}

func parseInput(filename string) ([]Equation, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([]Equation, len(lines))
	for i, line := range lines {
		splits := strings.Split(line, ":")
		nums, err := util.DelimitedStringOfNumbersToIntSlice(splits[1])
		if err != nil {
			return nil, err
		}
		result[i] = Equation{
			Value:   util.StringToNumber(splits[0]),
			Numbers: nums,
		}
	}

	return result, nil
}
