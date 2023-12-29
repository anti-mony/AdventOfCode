package main

import "fmt"

type q14store struct {
	board [][]string
}

func q14sol() error {

	inp, err := getInputQ14("input14.txt")
	if err != nil {
		return err
	}

	// rollNorth(inp)
	// fmt.Printf("P1 Answer is :%5d \n", totalLoadNorth(inp))

	fmt.Printf("P2 Answer is :%5d \n", solveQ14P2(inp))

	return nil
}

func solveQ14P2(inp [][]string) int {
	CYCLES := 1000000000

	ss := findSteadyState(inp)
	fmt.Println(ss)

	requiredCycles := CYCLES % (ss)
	if requiredCycles > ss {
		requiredCycles = requiredCycles - ss
	} else {
		requiredCycles = ss - requiredCycles
	}
	rollNCycles(inp, requiredCycles+1)
	return totalLoadNorth(inp)

}

func findSteadyState(inp [][]string) int {
	cache := []q14store{}
	for {
		rollNCycles(inp, 1)
		for s := 0; s < len(cache)-1; s++ {
			if areTheseTwoEqual(cache[s].board, inp) {
				return s + 1
			}
		}
		cache = append(cache, q14store{copy2DArray(inp)})
	}
}

func rollNCycles(board [][]string, cycles int) [][]string {
	for i := 0; i < cycles; i++ {
		rollNorth(board)
		rollWest(board)
		rollSouth(board)
		rollEast(board)
	}
	return board
}

func rollNorth(board [][]string) [][]string {

	nCols := len(board[0])
	nRows := len(board)

	for col := 0; col < nCols; col++ {
		for row := 1; row < nRows; row++ {
			if board[row][col] == "O" {
				current := row - 1
				for current >= 0 {
					if board[current][col] == "." {
						board[current][col] = "O"
						board[current+1][col] = "."
					} else if board[current][col] == "#" {
						break
					}
					current -= 1
				}
			}
		}
	}

	return board
}

func rollSouth(board [][]string) [][]string {
	nCols := len(board[0])
	nRows := len(board)

	for col := 0; col < nCols; col++ {
		for row := nRows - 1; row >= 0; row-- {
			if board[row][col] == "O" {
				current := row + 1
				for current < nRows {
					if board[current][col] == "." {
						board[current][col] = "O"
						board[current-1][col] = "."
					} else if board[current][col] == "#" {
						break
					}
					current += 1
				}
			}
		}
	}
	return board
}

func rollEast(board [][]string) [][]string {
	nCols := len(board[0])
	nRows := len(board)

	for row := 0; row < nRows; row++ {
		for col := nCols - 1; col >= 0; col-- {
			if board[row][col] == "O" {
				current := col + 1
				for current < nCols {
					if board[row][current] == "." {
						board[row][current] = "O"
						board[row][current-1] = "."
					} else if board[row][current] == "#" {
						break
					}
					current += 1
				}
			}
		}
	}

	return board
}

func rollWest(board [][]string) [][]string {
	nCols := len(board[0])
	nRows := len(board)

	for row := 0; row < nRows; row++ {
		for col := 1; col < nCols; col++ {
			if board[row][col] == "O" {
				current := col - 1
				for current >= 0 {
					if board[row][current] == "." {
						board[row][current] = "O"
						board[row][current+1] = "."
					} else if board[row][current] == "#" {
						break
					}
					current -= 1
				}
			}
		}
	}

	return board
}

func totalLoadNorth(inp [][]string) int {

	nCols := len(inp[0])
	nRows := len(inp)

	result := 0

	for row := 0; row < nRows; row++ {
		Os := 0
		for col := 0; col < nCols; col++ {
			if inp[row][col] == "O" {
				Os++
			}
		}
		result += Os * (nRows - row)
	}

	return result
}

func getInputQ14(filePath string) ([][]string, error) {

	lines, err := getFileAsListOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	result := make([][]string, len(lines))

	for i, line := range lines {
		l := make([]string, len(line))
		for j, c := range line {
			l[j] = string(c)
		}
		result[i] = l
	}

	return result, nil
}
