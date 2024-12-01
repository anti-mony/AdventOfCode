package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

const MAX_ENERGY = 9
const MIN_ENERGY = 0

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	inpQ2 := inp.Clone()

	fmt.Println("Answer Q1 ", Q1(inp, 100))
	fmt.Println("Answer Q2 ", Q2(inpQ2)+1)
}

func Q1(energies *grid.Grid[int], steps int) int {
	result := 0
	R, C := energies.Dimensions()

	for range steps {
		starts := make([]grid.Coordinate, 0)
		seen := make(map[grid.Coordinate]bool)
		for i := range R {
			for j := range C {
				c := grid.NewCoordinate(i, j)
				energies.IncrementValueAt(c, 1)
				if energies.ValueAt(c) > MAX_ENERGY {
					starts = append(starts, c)
				}
			}
		}
		for _, c := range starts {
			result += executeStep(energies, c, seen)
		}
		for i := range R {
			for j := range C {
				c := grid.NewCoordinate(i, j)
				if energies.ValueAt(c) > MAX_ENERGY {
					energies.SetValueAt(c, MIN_ENERGY)
				}
			}
		}
		// energies.Print()
	}

	return result
}

func Q2(energies *grid.Grid[int]) int {
	result := 0
	R, C := energies.Dimensions()
	step := 0
	for {
		starts := make([]grid.Coordinate, 0)
		seen := make(map[grid.Coordinate]bool)
		allNonZeros := true
		for i := range R {
			for j := range C {
				c := grid.NewCoordinate(i, j)
				energies.IncrementValueAt(c, 1)
				if energies.ValueAt(c) > MAX_ENERGY {
					starts = append(starts, c)
				}
				if energies.ValueAt(c) == 0 {
					allNonZeros = false
				}
			}
		}
		for _, c := range starts {
			result += executeStep(energies, c, seen)
		}
		allFlash := true
		for i := range R {
			for j := range C {
				c := grid.NewCoordinate(i, j)
				if energies.ValueAt(c) > MAX_ENERGY {
					energies.SetValueAt(c, MIN_ENERGY)
				} else {
					allFlash = false
				}
			}
		}
		if allFlash && allNonZeros {
			return step
		}
		step++
		// energies.Print()
	}
}

func executeStep(energies *grid.Grid[int], pos grid.Coordinate, seen map[grid.Coordinate]bool) int {
	if _, found := seen[pos]; found {
		return 0
	}
	q := list.NewQueue[grid.Coordinate]()
	flashes := 0
	seen[pos] = true
	q.Push(pos)
	// fmt.Println("------------------------")
	for q.Len() > 0 {
		c := q.Pop()
		flashes++

		for _, d := range grid.DIRECTIONS {
			n := c.Add(d)
			if energies.InBound(n) {
				energies.IncrementValueAt(n, 1)
				if _, found := seen[n]; !found {
					if energies.ValueAt(n) > MAX_ENERGY {
						seen[n] = true
						q.Push(n)
					}
				}
			}
		}
		// energies.Print()
	}
	// fmt.Println("------------------------")

	return flashes
}

func parseInput(filename string) (*grid.Grid[int], error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	return grid.NewIntGridFromStringSlice(lines)
}
