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

	fmt.Println("Answer P2 : ", solveP2(input))
}

func solveP1(rounds []Round) int {
	result := 0

	for _, round := range rounds {
		score := actionVal[round.Your]
		roundres := RockPaperScissor(round.Your, round.Opponent)
		score += resultVal[roundres]
		result += score
	}
	return result
}

func solveP2(rounds []Round) int {
	result := 0
	for _, round := range rounds {
		myMove := RockPaperScissorReverse(round.Opponent, round.Result)
		result += actionVal[myMove] + resultVal[round.Result]
	}
	return result
}

func RockPaperScissorReverse(opponent Action, result Result) Action {
	// Win
	if result == ResultWin {
		switch opponent {
		case ActionPaper:
			return ActionScissor
		case ActionScissor:
			return ActionRock
		case ActionRock:
			return ActionPaper
		}
	}

	if result == ResultLose {
		switch opponent {
		case ActionRock:
			return ActionScissor
		case ActionPaper:
			return ActionRock
		case ActionScissor:
			return ActionPaper
		}
	}

	return opponent
}

func RockPaperScissor(my, opponent Action) Result {
	if my == ActionRock {
		if opponent == ActionPaper {
			return ResultLose
		}
		if opponent == ActionScissor {
			return ResultWin
		}
	}

	if my == ActionPaper {
		if opponent == ActionScissor {
			return ResultLose
		}
		if opponent == ActionRock {
			return ResultWin
		}
	}

	if my == ActionScissor {
		if opponent == ActionRock {
			return ResultLose
		}
		if opponent == ActionPaper {
			return ResultWin
		}
	}

	return ResultDraw
}

type Action int

const (
	ActionRock Action = iota
	ActionPaper
	ActionScissor
)

type Result int

const (
	ResultWin Result = iota
	ResultLose
	ResultDraw
)

type Round struct {
	Opponent Action
	Your     Action
	Result   Result
}

var actionVal = map[Action]int{
	ActionScissor: 3,
	ActionPaper:   2,
	ActionRock:    1,
}

var resultVal = map[Result]int{
	ResultWin:  6,
	ResultDraw: 3,
	ResultLose: 0,
}

func parseInput(fileName string) ([]Round, error) {
	lines, err := util.GetFileAsListOfStrings(fileName)
	if err != nil {
		return nil, err
	}

	res := make([]Round, 0)
	for _, line := range lines {
		inp := strings.Split(line, " ")
		opponentAction, _ := getActionFromString(inp[0])
		myAction, result := getActionFromString(inp[1])
		res = append(res, Round{opponentAction, myAction, result})
	}

	return res, nil
}

func getActionFromString(in string) (Action, Result) {
	if in == "A" || in == "X" {
		return ActionRock, ResultLose
	}

	if in == "B" || in == "Y" {
		return ActionPaper, ResultDraw
	}

	return ActionScissor, ResultWin
}
