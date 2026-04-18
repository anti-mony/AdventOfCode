package util

import (
	"fmt"
)

// PrintMatrix prints a 2D array
func PrintMatrix[T any](in [][]T) {
	fmt.Println()
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0]); j++ {
			fmt.Printf("%3v", in[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// PrintMatrix prints a 2D array
func PrintMatrixWithIndices[T any](in [][]T) {
	fmt.Println()
	for i := -1; i < len(in); i++ {
		if i == -1 {
			for j := -1; j < len(in[0]); j++ {
				if j == -1 {
					fmt.Printf("%3v", "X")
					continue
				}
				fmt.Printf("%3v", j)
			}
			fmt.Println()
			continue
		}
		for j := -1; j < len(in[i]); j++ {
			if j == -1 {
				fmt.Printf("%3v", i)
				continue
			}
			fmt.Printf("%3v", in[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// CopyMatrix copies a 2D array
func CopyMatrix[T any](in [][]T) [][]T {
	result := make([][]T, len(in))
	for i := 0; i < len(in); i++ {
		row := make([]T, len(in[i]))
		copy(row, in[i])
		result[i] = row
	}
	return result
}

func MakeMatrix[T any](rows, cols int) [][]T {
	result := make([][]T, rows)
	for i := 0; i < rows; i++ {
		row := make([]T, cols)
		result[i] = row
	}
	return result
}

// AreEqual compares two matrices and returns a bool
func AreEqual[T comparable](a [][]T, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func ReadStringMatrixFromFile(filenme string) ([][]string, error) {
	lines, err := GetFileAsListOfStrings(filenme)
	if err != nil {
		return nil, err
	}

	result := make([][]string, len(lines))

	for i, line := range lines {
		result[i] = StringToCharSlice(line)
	}

	return result, nil
}

func ReadIntMatrixFromFile(filenme string) ([][]int, error) {
	lines, err := GetFileAsListOfStrings(filenme)
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(lines))

	for i, line := range lines {
		nums, err := StringOfNumbersToIntSlice(line)
		if err != nil {
			return nil, err
		}
		result[i] = nums
	}

	return result, nil
}

func FindIndexMatrix[T comparable](grid [][]T, v T) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == v {
				return i, j
			}
		}
	}
	return -1, -1
}

func Rotate[T any](grid [][]T, degree int) [][]T {
	grid = CopyMatrix(grid)
	switch degree {
	case 90:
		return rotate90(grid)
	case 180:
		return rotate90(rotate90(grid))
	case 270:
		return rotate90(rotate90(rotate90(grid)))
	}
	return grid
}

func rotate90[T any](grid [][]T) [][]T {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < i; j++ {
			grid[i][j], grid[j][i] = grid[j][i], grid[i][j]
		}
	}

	l, r := 0, len(grid[0])-1
	for l < r {
		for i := range len(grid) {
			grid[i][l], grid[i][r] = grid[i][r], grid[i][l]
		}
		l += 1
		r -= 1
	}
	return grid
}

func FlipVertically[T any](grid [][]T) [][]T {
	grid = CopyMatrix(grid)
	l, r := 0, len(grid[0])-1
	for l < r {
		for i := range len(grid) {
			grid[i][l], grid[i][r] = grid[i][r], grid[i][l]
		}
		l += 1
		r -= 1
	}
	return grid
}

func FlipHorizontally[T any](grid [][]T) [][]T {
	grid = CopyMatrix(grid)
	for i := range len(grid) / 2 {
		grid[i], grid[len(grid)-1-i] = grid[len(grid)-1-i], grid[i]
	}
	return grid
}
