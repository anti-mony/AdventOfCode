package main

import (
	"fmt"
	"log"

	"advent.of.code/grid"
	"advent.of.code/util"
)

const boardSize = 5

func main() {
	boards, randoms, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer P1: ", findRightBoard(boards, randoms))

	fmt.Println("Answer P2: ", findLastBoard(boards, randoms))
}

func findRightBoard(boards []*grid.Grid[int], randoms []int) int {
	for i, r := range randoms {
		for _, board := range boards {
			found := board.Find(r)
			if found != nil {
				board.SetValueAt(*found, -1)
			}
		}
		if i >= 5 {
			for _, board := range boards {
				if checkBingo(board) {
					return sumBoard(board) * r
				}
			}
		}
	}

	return -1
}

func findLastBoard(boards []*grid.Grid[int], randoms []int) int {
	winMap := make([]int, len(boards))
	windex := 1
	for _, r := range randoms {
		for _, board := range boards {
			found := board.Find(r)
			if found != nil {
				board.SetValueAt(*found, -1)
			}
		}
		for idx, board := range boards {
			if winMap[idx] == 0 {
				if checkBingo(board) {
					winMap[idx] = windex
					windex++
					if windex == len(boards)+1 {
						return sumBoard(board) * r
					}
				}
			}
		}
	}

	return -1
}

func sumBoard(board *grid.Grid[int]) int {
	L, W := board.Dimensions()
	total := 0
	for i := 0; i < L; i++ {
		for j := 0; j < W; j++ {
			v := board.ValueAt(grid.NewCoordinate(i, j))
			if v >= 0 {
				total += v
			}
		}
	}
	return total
}

func checkBingo(board *grid.Grid[int]) bool {

	L, W := board.Dimensions()
	// Check by ROW
	for i := 0; i < L; i++ {
		for j := 0; j < W; j++ {
			if board.ValueAt(grid.NewCoordinate(i, j)) != -1 {
				break
			}
			if j == W-1 {
				return true
			}
		}
	}

	// Check by COLS
	for j := 0; j < W; j++ {
		for i := 0; i < L; i++ {
			if board.ValueAt(grid.NewCoordinate(i, j)) != -1 {
				break
			}
			if i == L-1 {
				return true
			}
		}
	}

	return false
}

func parseInput(filename string) ([]*grid.Grid[int], []int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}

	randomInput, err := util.DelimitedStringOfNumbersToIntSlice(lines[0])
	if err != nil {
		return nil, nil, err
	}

	boards := make([]*grid.Grid[int], 0)
	nBoard := 0
	for i := 2; i+nBoard < len(lines); i += 5 {
		board, err := grid.NewIntGridFromDelimitedStringSlice(lines[i+nBoard : i+boardSize+nBoard])
		if err != nil {
			return nil, nil, err
		}
		boards = append(boards, board)
		nBoard++
	}

	return boards, randomInput, nil
}
