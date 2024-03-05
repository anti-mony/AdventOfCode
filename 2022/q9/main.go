package main

import (
	"fmt"
	"log"
	"strings"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {
	actions, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("> Answer P1: %d \n", solveP1(actions))

	fmt.Printf("> Answer P2: %d \n", solveP2(actions))
}

func solveP1(actions []Action) int {
	seen := make(map[grid.Coordinate]bool)
	H := grid.NewCoordinate(0, 0)
	T := grid.NewCoordinate(0, 0)

	for _, action := range actions {
		for i := 0; i < action.NumberOfSteps; i++ {
			newH := H.MoveTowards(action.Direction)
			if newH.DistanceFrom(T) > 1 {
				// fmt.Printf("newH: %v\t", newH)
				// fmt.Printf("T: %v\t", T)
				// fmt.Printf("newH to T: %d\t", newH.DistanceFrom(T))

				T = H

				// fmt.Printf("New T: %v\n", T)
				seen[T] = true
			}
			H = newH
		}
		// fmt.Println("-------------------------------------------------")
	}

	return len(seen) + 1
}

func solveP2([]Action) int {
	result := 0

	return result
}

type Action struct {
	Direction     grid.Direction
	NumberOfSteps int
}

func parseInput(filename string) ([]Action, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([]Action, 0)

	for _, line := range lines {
		rr := strings.Split(line, " ")
		result = append(result, Action{
			Direction:     grid.DirectionFromRLUD(rr[0]),
			NumberOfSteps: util.StringToNumber(rr[1]),
		})
	}

	return result, nil
}
