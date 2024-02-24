package main

import (
	"fmt"
	"log"
	"strings"

	"advent.of.code/util"
)

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer P1 : ", solveP1(input))
}

func solveP1(rounds []Round) int {
	result := 0

	for _, round := range rounds {
		score := actionVal[round.Your]
		roundres := RockPaperScissor(round.Your, round.Opponent)
		if roundres > 0 {
			score += 6
		} else if roundres < 0 {
			score += 0
		} else {
			score += 3
		}
		result += score
	}
	return result
}

func RockPaperScissor(my, opponent Action) int {
	if my == ActionRock {
		if opponent == ActionPaper {
			return -1
		}
		if opponent == ActionScissor {
			return 1
		}
	}

	if my == ActionPaper {
		if opponent == ActionScissor {
			return -1
		}
		if opponent == ActionRock {
			return 1
		}
	}

	if my == ActionScissor {
		if opponent == ActionRock {
			return -1
		}
		if opponent == ActionPaper {
			return 1
		}
	}

	return 0
}

type Action int

const (
	ActionRock Action = iota
	ActionPaper
	ActionScissor
)

type Round struct {
	Opponent Action
	Your     Action
}

var actionVal = map[Action]int{
	ActionScissor: 3,
	ActionPaper:   2,
	ActionRock:    1,
}

func parseInput(fileName string) ([]Round, error) {
	lines, err := util.GetFileAsListOfStrings(fileName)
	if err != nil {
		return nil, err
	}

	res := make([]Round, 0)
	for _, line := range lines {
		inp := strings.Split(line, " ")
		opponentAction := getActionFromString(inp[0])
		myAction := getActionFromString(inp[1])
		res = append(res, Round{opponentAction, myAction})
	}

	return res, nil
}

func getActionFromString(in string) Action {
	if in == "A" || in == "X" {
		return ActionRock
	}

	if in == "B" || in == "Y" {
		return ActionPaper
	}

	return ActionScissor
}
