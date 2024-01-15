package grid

import "fmt"

type Direction int

const (
	DirectionNorth Direction = iota + 1
	DirectionEast
	DirectionWest
	DirectionSouth
)

var (
	_north = Coordinate{-1, 0}
	_east  = Coordinate{0, 1}
	_south = Coordinate{1, 0}
	_west  = Coordinate{0, -1}
)

var DIRECTIONS = map[Direction]Coordinate{
	DirectionNorth: _north,
	DirectionEast:  _east,
	DirectionSouth: _south,
	DirectionWest:  _west,
}

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Add(i Coordinate) Coordinate {
	return Coordinate{c.x + i.x, c.y + i.y}
}

// PrintGrid prints a 2D array
func PrintGrid[T string | int](in [][]T) {
	fmt.Println()
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			fmt.Printf("%2v", in[i][j])
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
