package grid

import "fmt"

//go:generate stringer -type=Direction
type Direction int

const (
	DirectionNorth Direction = iota + 1
	DirectionEast
	DirectionWest
	DirectionSouth
)

func (d Direction) Reverse() Direction {
	switch d {
	case DirectionEast:
		return DirectionWest
	case DirectionNorth:
		return DirectionSouth
	case DirectionSouth:
		return DirectionNorth
	}
	return DirectionEast
}

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
	X int
	Y int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{X: x, Y: y}
}

func (c Coordinate) Add(i Coordinate) Coordinate {
	return Coordinate{c.X + i.X, c.Y + i.Y}
}

func (c Coordinate) MoveTowards(d Direction) Coordinate {
	delta := DIRECTIONS[d]
	return Coordinate{c.X + delta.X, c.Y + delta.Y}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

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
