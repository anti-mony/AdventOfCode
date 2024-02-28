package main

import (
	"fmt"
	"log"
	"regexp"

	"advent.of.code/list"
	"advent.of.code/util"
)

const (
	stackVarWidth = 3 // eg. [2]
)

func main() {
	stacks, actions, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Answer P1: ", SolveP1(stacks, actions))

	fmt.Println("Answer P2: ", SolveP2(stacks, actions))
}

func SolveP1(stacks []*list.Stack[string], actions []Action) string {
	result := ""

	for _, action := range actions {
		for i := 0; i < action.ContainersToMove; i++ {
			stacks[action.ToStackIndex].Push(stacks[action.FromStackIndex].Pop())
		}
	}

	for _, s := range stacks {
		result += s.Peek()
	}

	return result
}

func SolveP2(stacks []*list.Stack[string], actions []Action) string {
	result := ""

	for _, action := range actions {
		tmp := list.NewStack[string]()

		for i := 0; i < action.ContainersToMove; i++ {
			tmp.Push(stacks[action.FromStackIndex].Pop())
		}

		for i := 0; i < action.ContainersToMove; i++ {
			stacks[action.ToStackIndex].Push(tmp.Pop())
		}
	}

	for _, s := range stacks {
		result += s.Peek()
	}

	return result
}

type Action struct {
	FromStackIndex   int
	ToStackIndex     int
	ContainersToMove int
}

func parseInput(fileName string) ([]*list.Stack[string], []Action, error) {
	lines, err := util.GetFileAsListOfStrings(fileName)
	if err != nil {
		return nil, nil, err
	}

	// Assuming Square box input
	N := len(lines[0])

	numberOfStacks := (N + 1) / 4

	stacks := make([]*list.Stack[string], numberOfStacks)

	for i := 0; i < numberOfStacks; i++ {
		stacks[i] = list.NewStack[string]()
	}

	lineNo := findEmptyLineIndex(lines)

	for i := lineNo - 2; i >= 0; i-- {
		line := lines[i]
		for i := 0; i < N; i += 4 {
			stackN := (i + 1) / (stackVarWidth + 1)
			if string(line[i+1]) != " " {
				stacks[stackN].Push(string(line[i+1]))
			}
		}
	}

	actions := make([]Action, 0)

	re := regexp.MustCompile(`(?m)\d+`)
	for _, line := range lines[lineNo+1:] {
		matches := re.FindAllString(line, -1)
		actions = append(actions, Action{
			FromStackIndex:   util.StringToNumber(matches[1]) - 1,
			ToStackIndex:     util.StringToNumber(matches[2]) - 1,
			ContainersToMove: util.StringToNumber(matches[0]),
		})
	}

	return stacks, actions, nil
}

func findEmptyLineIndex(in []string) int {
	for idx, l := range in {
		if len(l) == 0 {
			return idx
		}
	}
	return -1
}
