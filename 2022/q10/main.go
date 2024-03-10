package main

import (
	"fmt"
	"log"
	"strings"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {
	commands, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("> Answer P1: %d \n", solveP1(commands))

	solveP2(commands)
}

func solveP2(commands []Command) {
	register := 1
	current := 0
	tick := 0

	CRT_Rows := 6
	CRT_Cols := 40

	CRT := make([][]string, CRT_Rows)

	for i := 0; i < CRT_Rows; i++ {
		CRT[i] = make([]string, CRT_Cols)
	}

	for i := 0; i < 240; i++ {
		// Draw
		spritePos := register % CRT_Cols
		CRTRow := i / CRT_Cols
		CRTCol := i - (CRTRow * CRT_Cols)

		if CRTCol == spritePos-1 || CRTCol == spritePos || CRTCol == spritePos+1 {
			CRT[CRTRow][CRTCol] = "#"
		} else {
			CRT[CRTRow][CRTCol] = "."
		}

		if commands[current].NoOp {
			current++
		} else if commands[current].Add {
			tick++
			if tick == 2 {
				register += commands[current].Value
				tick = 0
				current++
			}
		}
	}

	grid.PrintGrid(CRT)
}

func solveP1(commands []Command) int {
	cycles := map[int]int{
		20:  1,
		60:  1,
		100: 1,
		140: 1,
		180: 1,
		220: 1,
	}

	register := 1
	current := 0
	tick := 0

	for i := 1; i < 221; i++ {

		if _, ok := cycles[i]; ok {
			cycles[i] = register
		}

		if commands[current].NoOp {
			current++
		} else if commands[current].Add {
			tick++
			if tick == 2 {
				register += commands[current].Value
				tick = 0
				current++
			}
		}

	}

	result := 0
	for k, v := range cycles {
		result += k * v
	}

	return result
}

type Command struct {
	NoOp  bool
	Add   bool
	Value int
}

func parseInput(filename string) ([]Command, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	commands := make([]Command, 0)

	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			commands = append(commands, Command{NoOp: true})
		} else {
			splits := strings.Split(line, " ")
			commands = append(commands, Command{Add: true, Value: util.StringToNumber(splits[1])})
		}
	}

	return commands, nil
}
