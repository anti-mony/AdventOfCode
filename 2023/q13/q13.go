package main

import (
	"fmt"
	"slices"
)

func q13sol() error {

	input, err := getInputQ13("input13.txt")
	if err != nil {
		return err
	}

	fmt.Printf("P1 Answer is : %d\n", getResForQ13_P1(input))

	fmt.Printf("P2 Answer is : %d\n", getResForQ13_P2(input))

	return nil
}

func getResForQ13_P1(input [][][]string) int {

	rows, cols := 0, 0
	for _, p := range input {
		// fmt.Println()
		// for i := 0; i < len(p); i++ {
		// 	for j := 0; j < len(p[i]); j++ {
		// 		fmt.Printf("%2s", p[i][j])
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()
		r := findHorizontalLine(p)
		rows += r
		c := findVerticalLine(p)
		cols += c

		// fmt.Printf("ROWS: %d, COLS:%d \n", r, c)
	}

	return rows*100 + cols
}

func getResForQ13_P2(input [][][]string) int {

	rows, cols := 0, 0
	for idx, p := range input {

		// fmt.Println()
		// for i := 0; i < len(p); i++ {
		// 	for j := 0; j < len(p[i]); j++ {
		// 		fmt.Printf("%2s", p[i][j])
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()

		hWorks, vWorks := make([]int, 1), make([]int, 1)

		worked := false

		for r := 0; r < len(p); r++ {
			worked = doesItWorkHorizontally(p, r, r+1)
			if worked {
				hWorks = append(hWorks, r+1)
				break
			}
		}

		worked = false
		for c := 0; c < len(p[0]); c++ {
			worked = doesItWorkVertically(p, c, c+1)
			if worked {
				vWorks = append(vWorks, c+1)
				break
			}
		}

		fmt.Printf("%2d) ROWS: %v, COLS:%v \n", idx, hWorks, vWorks)

		rows += slices.Max(hWorks)
		cols += slices.Max(vWorks)

	}

	return rows*100 + cols
}

func findVerticalLine(input [][]string) int {
	doesItWork := func(p [][]string, left, right int) bool {
		rows := len(p)
		left -= 1
		right += 1
		for left >= 0 && right < len(p[0]) {
			r := 0
			for r < rows {
				if p[r][left] != p[r][right] {
					return false
				}
				r += 1
			}
			left -= 1
			right += 1
		}
		return true
	}

	for col := 0; col < len(input[0])-1; col++ {
		match := true
		row := 0
		for row = 0; row < len(input); row++ {
			if input[row][col] != input[row][col+1] {
				match = false
				break
			}
		}
		if match {
			if doesItWork(input, col, col+1) {
				return col + 1
			}
		}
	}

	return 0
}

func findHorizontalLine(input [][]string) int {

	doesItWork := func(p [][]string, top, bottom int) bool {
		cols := len(p[0])
		top -= 1
		bottom += 1
		for top >= 0 && bottom < len(p) {
			c := 0
			for c < cols {
				if p[top][c] != p[bottom][c] {
					return false
				}
				c += 1
			}
			top -= 1
			bottom += 1
		}
		return true
	}

	for row := 0; row < len(input)-1; row++ {
		match := true
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] != input[row+1][col] {
				match = false
				break
			}
		}
		if match {
			if doesItWork(input, row, row+1) {
				return row + 1
			}
		}
	}

	return 0
}

func doesItWorkVertically(p [][]string, left, right int) bool {
	mismatches := 0
	rows := len(p)
	for left >= 0 && right < len(p[0]) {
		r := 0
		for r < rows {
			if p[r][left] != p[r][right] {
				mismatches += 1
				if mismatches > 1 {
					return false
				}
			}
			r += 1
		}
		left -= 1
		right += 1
	}
	return mismatches == 1
}

func doesItWorkHorizontally(p [][]string, top, bottom int) bool {
	mismatches := 0

	cols := len(p[0])
	for top >= 0 && bottom < len(p) {
		c := 0
		for c < cols {
			if p[top][c] != p[bottom][c] {
				mismatches += 1
				if mismatches > 1 {
					return false
				}
			}
			c += 1
		}
		top -= 1
		bottom += 1
	}
	return mismatches == 1
}

func getInputQ13(filePath string) ([][][]string, error) {

	lines, err := getFileAsListOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	result := make([][][]string, 0)

	puzzle := make([][]string, 0)
	for _, line := range lines {
		if line == "" {
			result = append(result, puzzle)
			puzzle = make([][]string, 0)
			continue
		}
		l := make([]string, 0)
		for _, c := range line {
			l = append(l, string(c))
		}
		puzzle = append(puzzle, l)
	}
	result = append(result, puzzle)

	return result, nil
}
