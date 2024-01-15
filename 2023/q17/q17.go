package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

type qVar struct {
	current          grid.Coordinate
	direction        grid.Direction
	stepsInDirection int
	heatLoss         int
}

type seenVar struct {
	current          grid.Coordinate
	direction        grid.Direction
	stepsInDirection int
}

func isLessThan(a, b any) bool {
	aA := a.(qVar)
	bB := b.(qVar)

	return aA.heatLoss < bB.heatLoss
}

func main() {
	fileName := os.Args[1]
	g, err := getInput(fileName)
	if err != nil {
		log.Fatal(err, g)
	}

	// grid.PrintGrid(g)

	fmt.Printf("Answer 1 solution is %d\n", solveP1(g))

}

func solveP1(in [][]int) int {
	result := 0

	seen := make(map[seenVar]bool)

	pq := list.NewHeap(isLessThan)

	pq.Push(qVar{
		grid.Coordinate{X: 0, Y: 0}, 0, 0, 0,
	})

	for !pq.IsEmpty() {
		vAny := pq.Pop()
		v := vAny.(qVar)

		if v.current.X == len(in)-1 && v.current.Y == len(in[0])-1 {
			return v.heatLoss
		}

		if _, ok := seen[seenVar{v.current, v.direction, v.stepsInDirection}]; ok {
			continue
		}
		seen[seenVar{v.current, v.direction, v.stepsInDirection}] = true

		if v.stepsInDirection < 3 && v.direction != 0 {
			delta := grid.DIRECTIONS[v.direction]
			new := v.current.Add(delta)

			if new.X >= 0 && new.X < len(in) && new.Y >= 0 && new.Y < len(in[0]) {
				pq.Push(qVar{
					new, v.direction, v.stepsInDirection + 1, v.heatLoss + in[new.X][new.Y],
				})
			}
		}

		for i := 1; i <= 4; i++ {
			if v.direction != grid.Direction(i) && v.direction.Reverse() != grid.Direction(i) {
				delta := grid.DIRECTIONS[grid.Direction(i)]
				new := v.current.Add(delta)

				if new.X >= 0 && new.X < len(in) && new.Y >= 0 && new.Y < len(in[0]) {
					pq.Push(qVar{
						new, grid.Direction(i), 1, v.heatLoss + in[new.X][new.Y],
					})
				}
			}
		}
	}

	return result
}

func getInput(filename string) ([][]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(lines))

	for i, line := range lines {
		row := make([]int, len(line))
		for ii, c := range line {
			row[ii] = util.StringToNumber(string(c))
		}
		result[i] = row
	}

	return result, nil

}
