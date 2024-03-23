package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	fileName := "input.small.txt"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	input, err := parseInput(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("> Answer P1: ", solveP1(input))
}

type qVar struct {
	Cd   grid.Coordinate
	Dist int
}

func (q qVar) String() string {
	return fmt.Sprintf("[%v, %d]", q.Cd, q.Dist)
}

func solveP1(in PuzzleInput) int {
	q := list.NewQueue[qVar]()
	q.Push(qVar{in.Start, 0})

	seen := map[grid.Coordinate]any{in.Start: struct{}{}}

	for q.Len() > 0 {
		current := q.Pop()

		for i := 1; i < 5; i++ {
			dir := grid.Direction(i)
			next := current.Cd.MoveTowards(dir)
			_, ok := seen[next]
			if !ok && in.Mt.InBound(next) {
				if in.Mt.ValueAt(next) <= in.Mt.ValueAt(current.Cd)+1 {
					if next == in.End {
						return current.Dist + 1
					}
					seen[next] = struct{}{}
					q.Push(qVar{next, current.Dist + 1})
				}
			}
		}

	}

	return -1
}

type PuzzleInput struct {
	Mountain [][]rune
	Start    grid.Coordinate
	End      grid.Coordinate
	Mt       *grid.Grid[rune]
}

func parseInput(filename string) (PuzzleInput, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return PuzzleInput{}, err
	}

	rows := len(lines)
	cols := len(lines[0])
	mtGrid := grid.NewGrid[rune](rows, cols)

	mountainMap := make([][]rune, 0)
	start := grid.NewCoordinate(0, 0)
	end := grid.NewCoordinate(0, 0)

	for ln, line := range lines {
		c := make([]rune, len(line))
		for i, cc := range line {
			mtGrid.SetValueAt(grid.NewCoordinate(ln, i), cc)
			if string(cc) == "S" {
				start = grid.NewCoordinate(ln, i)
				mtGrid.SetValueAt(grid.NewCoordinate(ln, i), 'a')
			}
			if string(cc) == "E" {
				end = grid.NewCoordinate(ln, i)
				mtGrid.SetValueAt(grid.NewCoordinate(ln, i), 'z')
			}

			c[i] = cc
		}
		mountainMap = append(mountainMap, c)
	}

	return PuzzleInput{
		Mountain: mountainMap,
		Start:    start,
		End:      end,
		Mt:       mtGrid,
	}, nil
}
