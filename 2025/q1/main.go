package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/util"
)

const DIAL = 100

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Q1: %d\n", Q1(inp))
	fmt.Printf("Q2: %d\n", Q2(inp))
}

type Rotation struct {
	Value     int
	Direction int // 0 Left, 1 Right
}

func parseInput(filename string) ([]Rotation, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	inp := make([]Rotation, len(lines))
	for i, line := range lines {
		dir := 0
		if line[0] == 'R' {
			dir = 1
		}
		inp[i] = Rotation{
			Direction: dir,
			Value:     util.StringToNumber(line[1:]),
		}
	}

	return inp, nil
}

func Q1(inp []Rotation) int {
	result := 0
	current := 50
	for _, rot := range inp {
		if rot.Direction == 0 {
			current -= (rot.Value)
		} else {
			current += rot.Value
		}
		current = current % DIAL
		if current < 0 {
			current += DIAL
		} else if current == 100 {
			current = 0
		}
		if current == 0 {
			result++
		}
	}
	return result
}

func Q2(inp []Rotation) int {
	result := 0
	current := 50
	for _, rot := range inp {
		completeRotations := rot.Value / DIAL
		rotVal := rot.Value % DIAL
		newValue := current
		if rot.Direction == 0 {
			newValue -= rotVal
		} else {
			newValue += rotVal
		}

		fmt.Printf("> %4d %4d %4d %2d %4d \n", current, newValue, completeRotations, rot.Direction, rotVal)
		if (current > 0 && newValue < 0) ||
			newValue > DIAL ||
			rotVal == 0 {
			result += 1
			fmt.Println("Result Incremented because crossed 0", result)
		}

		result += completeRotations

		current = newValue
		current = current % DIAL
		if current < 0 {
			current += DIAL
		} else if current == 100 {
			current = 0
		}
		if current == 0 {
			result++
			fmt.Println("Result Incremented because exact 0", result)

		}
	}

	return result
}
