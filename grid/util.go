package grid

import "fmt"

// PrintGrid prints a 2D array
func PrintGrid[T string | int](in [][]T) {
	fmt.Println()
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			fmt.Printf("%v", in[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// CopyGrid copies a 2D array
func CopyGrid[T string | int](in [][]T) [][]T {
	result := make([][]T, len(in))
	for i := 0; i < len(in); i++ {
		row := make([]T, len(in[i]))
		copy(row, in[i])
		result[i] = row
	}
	return result
}

// AreEqual compares two grid and returns a bool
func AreEqual[T string | int](a [][]T, b [][]T) bool {
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
