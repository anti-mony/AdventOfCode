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

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solve Q1: ", solveQ1(inp))
	fmt.Println("Solve Q2: ", solveQ2(inp))
}

func solveQ1(lines [][]string) int {
	cost := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	pair := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	result := 0
	for _, line := range lines {
		stack := list.NewStack[string]()
		for _, c := range line {
			if c == "(" || c == "[" || c == "{" || c == "<" {
				stack.Push(c)
			} else {
				p := stack.Pop()
				if p != pair[c] {
					result += cost[c]
					break
				}
			}
		}
	}

	return result
}

func solveQ2(lines [][]string) int {
	cost := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	pair := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	scores := make([]int, 0)
	for _, line := range lines {
		stack := list.NewStack[string]()
		i := 0
		c := ""
		for i, c = range line {
			if c == "(" || c == "[" || c == "{" || c == "<" {
				stack.Push(c)
			} else {
				p := stack.Pop()
				if p != pair[c] {
					break
				}
			}
		}
		if i == len(line)-1 {
			score := 0
			for stack.Len() > 0 {
				score *= 5
				score += cost[stack.Pop()]
			}
			scores = append(scores, score)
		}
	}

	slices.Sort(scores)

	return scores[len(scores)/2]
}

func parseInput(filename string) ([][]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	res := make([][]string, len(lines))
	for i, line := range lines {
		res[i] = util.StringToCharSlice(line)
	}

	return res, nil
}
