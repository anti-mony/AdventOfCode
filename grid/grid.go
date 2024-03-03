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

type grid[T any] struct {
	Store [][]T
}

func NewGrid[T any](rows, columns int) *grid[T] {
	grid := &grid[T]{
		Store: make([][]T, rows),
	}
	for idx := range grid.Store {
		grid.Store[idx] = make([]T, columns)
	}

	return grid
}

func (g *grid[T]) Dimensions() (int, int) {
	rows, cols := 0, 0
	if g == nil {
		return rows, cols
	}

	rows = len(g.Store)
	if rows == 0 {
		return rows, cols
	}

	cols = len(g.Store[0])

	return rows, cols
}

func (g *grid[T]) InBound(c Coordinate) bool {
	return c.X >= 0 && c.X < len(g.Store) && c.Y >= 0 && c.Y < len(g.Store[0])
}

func (g *grid[T]) ValueAt(c Coordinate) T {
	return g.Store[c.X][c.Y]
}
