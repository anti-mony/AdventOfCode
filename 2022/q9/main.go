package main

import (
	"fmt"
	"log"
	"math"
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

func solveP2(actions []Action) int {
	seen := make(map[grid.Coordinate]bool)
	snake := make([]grid.Coordinate, 0)
	snakeLen := 10
	for i := 0; i < snakeLen; i++ {
		snake = append(snake, grid.NewCoordinate(0, 0))
	}

	for _, action := range actions {
		for i := 0; i < action.NumberOfSteps; i++ {
			snake[0] = snake[0].MoveTowards(action.Direction)
			for i := 0; i < snakeLen-1; i++ {
				dX := snake[i].X - snake[i+1].X
				dY := snake[i].Y - snake[i+1].Y
				if math.Abs(float64(dX)) > 1 || math.Abs(float64(dY)) > 1 {
					if dX == 0 {
						snake[i+1].Y += dY / 2
					} else if dY == 0 {
						snake[i+1].X += dX / 2
					} else {
						if dX > 0 {
							snake[i+1].X += 1
						} else {
							snake[i+1].X -= 1
						}
						if dY > 0 {
							snake[i+1].Y += 1
						} else {
							snake[i+1].Y -= 1
						}
					}
				}
			}
			seen[snake[snakeLen-1]] = true
		}
	}

	return len(seen)
}

func moveSnake(snake []grid.Coordinate, dir grid.Direction) {
	for i := 0; i < len(snake)-1; i++ {
		snake[i] = snake[i+1]
	}
	snake[len(snake)-1] = snake[len(snake)-2].MoveTowards(dir)
}

func solveP1(actions []Action) int {
	seen := make(map[grid.Coordinate]bool)
	H := grid.NewCoordinate(0, 0)
	T := grid.NewCoordinate(0, 0)

	for _, action := range actions {
		for i := 0; i < action.NumberOfSteps; i++ {
			newH := H.MoveTowards(action.Direction)
			if newH.DistanceFrom(T) > 1 {
				T = H
				seen[T] = true
			}
			H = newH
		}
	}

	return len(seen) + 1
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
