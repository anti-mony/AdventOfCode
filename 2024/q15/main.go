package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	warehouse, movements, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	warehouse2 := util.CopyMatrix(warehouse)

	// fmt.Println("A1: ", Q1(warehouse, movements))

	fmt.Println("A2: ", Q2(warehouse2, movements))
}

var movementDirection = map[string]grid.Direction{
	"^": grid.DirectionNorth,
	">": grid.DirectionEast,
	"v": grid.DirectionSouth,
	"<": grid.DirectionWest,
}

func Q1(warehouse [][]string, movements []string) int {

	ri, rj := util.FindIndexMatrix(warehouse, "@")
	r := grid.NewCoordinate(ri, rj)

	for _, mov := range movements {
		dir := movementDirection[mov]
		n := r.Add(grid.DIRECTIONS[dir])
		if !inBound(warehouse, n) {
			continue
		}
		if warehouse[n.X][n.Y] == "#" {
			continue
		}
		if warehouse[n.X][n.Y] == "." {
			warehouse[r.X][r.Y] = "."
			warehouse[n.X][n.Y] = "@"
			r = n
			continue
		}

		// find empty space
		for inBound(warehouse, n) {
			if warehouse[n.X][n.Y] == "." {
				break
			} else if warehouse[n.X][n.Y] == "#" {
				// there's no empty space we've hit a wall here
				break
			}
			n = n.Add(grid.DIRECTIONS[dir])
		}

		// if it's out of bounds or a wall, move to next movement
		if !inBound(warehouse, n) || warehouse[n.X][n.Y] == "#" {
			continue
		}

		// we need to move all the blocks ahead of it to the right position
		for n != r.Add(grid.DIRECTIONS[dir]) {
			warehouse[n.X][n.Y] = "O"
			n = n.Add(grid.DIRECTIONS[dir.Reverse()])
			warehouse[n.X][n.Y] = "."
		}
		warehouse[r.X][r.Y] = "."
		warehouse[n.X][n.Y] = "@"
		r = n
	}

	result := 0
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == "O" {
				result += i*100 + j
			}
		}
	}

	return result
}

func Q2(warehouse [][]string, movements []string) int {

	// Double the matrix first
	ri, rj := -1, -1
	newWarehouse := make([][]string, len(warehouse))
	for i := 0; i < len(warehouse); i++ {
		newWarehouse[i] = make([]string, len(warehouse[i])*2)
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == "@" {
				ri, rj = i, 2*j
				newWarehouse[i][2*j] = "@"
				newWarehouse[i][2*j+1] = "."
			} else if warehouse[i][j] == "O" {
				newWarehouse[i][2*j] = "["
				newWarehouse[i][2*j+1] = "]"
			} else {
				newWarehouse[i][2*j] = warehouse[i][j]
				newWarehouse[i][2*j+1] = warehouse[i][j]
			}
		}
	}
	warehouse = newWarehouse

	// Starting Position
	r := grid.NewCoordinate(ri, rj)

	for _, mov := range movements {
		// util.PrintMatrix(warehouse)
		dir := movementDirection[mov]
		n := r.Add(grid.DIRECTIONS[dir])
		// out of bound, move on
		if !inBound(warehouse, n) {
			continue
		}

		// wall ahead, move on
		if warehouse[n.X][n.Y] == "#" {
			continue
		}

		// clear to move move the robot and continue
		if warehouse[n.X][n.Y] == "." {
			warehouse[r.X][r.Y] = "."
			warehouse[n.X][n.Y] = "@"
			r = n
			continue
		}

		// Lets do by directions

		// East West
		if dir == grid.DirectionEast || dir == grid.DirectionWest {
			n := r.Add(grid.DIRECTIONS[dir])
			// find empty space
			for inBound(warehouse, n) {
				if warehouse[n.X][n.Y] == "." {
					break
				} else if warehouse[n.X][n.Y] == "#" {
					// there's no empty space we've hit a wall here
					break
				}
				n = n.Add(grid.DIRECTIONS[dir])
			}

			// if it's out of bounds or a wall, move to next movement
			if !inBound(warehouse, n) || warehouse[n.X][n.Y] == "#" {
				continue
			}

			// we need to move all the blocks ahead of it to the right position
			for n != r.Add(grid.DIRECTIONS[dir]) {
				t := n.Add(grid.DIRECTIONS[dir.Reverse()])
				warehouse[n.X][n.Y] = warehouse[t.X][t.Y]
				warehouse[t.X][t.Y] = "."
				n = t
			}
			warehouse[r.X][r.Y] = "."
			warehouse[n.X][n.Y] = "@"
			r = n
		} else {
			// North South

			blocksToMove := getBlocksToMove(warehouse, r, dir)

			topLayer := map[int][]grid.Coordinate{}
			topLayerList := []grid.Coordinate{}

			for i := len(blocksToMove) - 1; i >= 0; i-- {
				for j := 0; j < len(blocksToMove[i]); j++ {
					topLayer[blocksToMove[i][j].Y] = append(topLayer[blocksToMove[i][j].Y], blocksToMove[i][j])
				}
			}

			for k := range topLayer {
				slices.SortFunc(topLayer[k], func(a, b grid.Coordinate) int {
					if dir == grid.DirectionSouth {
						return a.X - b.X
					}
					return b.X - a.X
				})
				topLayerList = append(topLayerList, topLayer[k][len(topLayer[k])-1])
			}

			// find empty space next to the last layer
			canMoveBlocks := true
			for _, b := range topLayerList {
				bn := b.Add(grid.DIRECTIONS[dir])
				if !inBound(warehouse, bn) || warehouse[bn.X][bn.Y] != "." {
					canMoveBlocks = false
				}
			}

			if !canMoveBlocks {
				continue
			}

			// if we're here that means we can move the whole block one layer up
			for i := len(blocksToMove) - 1; i >= 0; i-- {
				for j := 0; j < len(blocksToMove[i]); j++ {
					n := blocksToMove[i][j].Add(grid.DIRECTIONS[dir])
					warehouse[n.X][n.Y] = warehouse[blocksToMove[i][j].X][blocksToMove[i][j].Y]
					warehouse[blocksToMove[i][j].X][blocksToMove[i][j].Y] = "."
				}
			}
			n := r.Add(grid.DIRECTIONS[dir])
			warehouse[n.X][n.Y] = "@"
			warehouse[r.X][r.Y] = "."
			r = n
		}
	}

	util.PrintMatrix(warehouse)

	result := 0
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == "[" {
				result += i*100 + j
			}
		}
	}

	return result
}

func getBlocksToMove(warehouse [][]string, robot grid.Coordinate, direction grid.Direction) [][]grid.Coordinate {
	result := [][]grid.Coordinate{}
	delta := grid.DIRECTIONS[direction]
	q := list.NewQueue[grid.Coordinate]()
	n := robot.Add(delta)
	if inBound(warehouse, n) {
		if warehouse[n.X][n.Y] == "[" {
			q.Push(n)
			q.Push(grid.NewCoordinate(n.X, n.Y+1))
		} else {
			q.Push(n)
			q.Push(grid.NewCoordinate(n.X, n.Y-1))
		}
	}

	for q.Len() > 0 {
		N := q.Len()
		layer := []grid.Coordinate{}
		for range N {
			current := q.Pop()
			layer = append(layer, current)
			n := current.Add(delta)
			currentVal := warehouse[current.X][current.Y]
			if inBound(warehouse, n) {
				if currentVal == "[" && warehouse[n.X][n.Y] == "[" {
					q.Push(n)
				} else if currentVal == "]" && warehouse[n.X][n.Y] == "]" {
					q.Push(n)
				} else if currentVal == "[" && warehouse[n.X][n.Y] == "]" {
					q.Push(n)
					q.Push(grid.NewCoordinate(n.X, n.Y-1))
				} else if currentVal == "]" && warehouse[n.X][n.Y] == "[" {
					q.Push(n)
					q.Push(grid.NewCoordinate(n.X, n.Y+1))
				}
			}
		}
		slices.SortFunc(layer, func(a, b grid.Coordinate) int { return a.Y - b.Y })
		result = append(result, layer)
	}

	return result
}

func inBound(warehouse [][]string, c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(warehouse) && c.Y >= 0 && c.Y < len(warehouse[c.X])
}

func parseInput(filename string) ([][]string, []string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}

	movements := []string{}
	warehouse := [][]string{}
	i := 0

	for i = 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}
		warehouse = append(warehouse, util.StringToCharSlice(lines[i]))
	}

	i++
	for ; i < len(lines); i++ {
		movements = append(movements, util.StringToCharSlice(lines[i])...)
	}

	return warehouse, movements, nil
}
