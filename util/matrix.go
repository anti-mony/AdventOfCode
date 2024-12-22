package util

import (
	"fmt"
)

// PrintMatrix prints a 2D array
func PrintMatrix[T any](in [][]T) {
	fmt.Println()
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			fmt.Printf("%3v", in[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// CopyMatrix copies a 2D array
func CopyMatrix[T comparable](in [][]T) [][]T {
	result := make([][]T, len(in))
	for i := 0; i < len(in); i++ {
		row := make([]T, len(in[i]))
		copy(row, in[i])
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
