package grid

import (
	"cmp"
	"fmt"
	"math"
	"strings"

	"advent.of.code/util"
)

//go:generate stringer -type=Direction
type Direction int

const (
	DirectionNorth Direction = iota + 1
	DirectionEast
	DirectionWest
	DirectionSouth
	DirectionNorthEast
	DirectionNorthWest
	DirectionSouthEast
	DirectionSouthWest
)

func (d Direction) Reverse() Direction {
	switch d {
	case DirectionEast:
		return DirectionWest
	case DirectionNorth:
		return DirectionSouth
	case DirectionSouth:
		return DirectionNorth
	case DirectionNorthEast:
		return DirectionSouthWest
	case DirectionNorthWest:
		return DirectionSouthEast
	case DirectionSouthEast:
		return DirectionNorthWest
	case DirectionSouthWest:
		return DirectionNorthEast
	}
	return DirectionEast
}

func DirectionFromRLUD(in string) Direction {
	switch strings.ToUpper(in) {
	case "R":
		return DirectionEast
	case "L":
		return DirectionWest
	case "U":
		return DirectionNorth
	}

	return DirectionSouth
}

var (
	_north     = Coordinate{-1, 0}
	_east      = Coordinate{0, 1}
	_south     = Coordinate{1, 0}
	_west      = Coordinate{0, -1}
	_northEast = Coordinate{-1, 1}
	_northWest = Coordinate{-1, -1}
	_southEast = Coordinate{1, 1}
	_southWest = Coordinate{1, -1}
)

var DIRECTIONS = map[Direction]Coordinate{
	DirectionNorth:     _north,
	DirectionEast:      _east,
	DirectionSouth:     _south,
	DirectionWest:      _west,
	DirectionNorthEast: _northEast,
	DirectionNorthWest: _northWest,
	DirectionSouthEast: _southEast,
	DirectionSouthWest: _southWest,
}

var DIRECTIONS_STRAIGHT = map[Direction]Coordinate{
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

func (c Coordinate) MoveNUnitsTowards(d Direction, value int) Coordinate {
	delta := DIRECTIONS[d]
	return Coordinate{c.X + delta.X*value, c.Y + delta.Y*value}
}

func (c Coordinate) DistanceFrom(d Coordinate) int {
	x2 := (c.X - d.X) * (c.X - d.X)
	y2 := (c.Y - d.Y) * (c.Y - d.Y)

	return int(math.Sqrt(float64(x2) + float64(y2)))
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func (c Coordinate) Ptr() *Coordinate {
	return &c
}

type Grid[T cmp.Ordered] struct {
	store [][]T
}

func NewGrid[T cmp.Ordered](rows, columns int) *Grid[T] {
	grid := &Grid[T]{
		store: make([][]T, rows),
	}
	for idx := range grid.store {
		grid.store[idx] = make([]T, columns)
	}

	return grid
}

func NewIntGridFromDelimitedStringSlice(input []string) (*Grid[int], error) {
	rows := len(input)
	grid := &Grid[int]{
		store: make([][]int, rows),
	}
	var err error
	for idx := range grid.store {
		grid.store[idx], err = util.DelimitedStringOfNumbersToIntSlice(input[idx])
		if err != nil {
			return nil, err
		}
	}
	return grid, nil
}

func NewIntGridFromStringSlice(input []string) (*Grid[int], error) {
	rows := len(input)
	grid := &Grid[int]{
		store: make([][]int, rows),
	}
	var err error
	for idx := range grid.store {
		grid.store[idx], err = util.StringOfNumbersToIntSlice(input[idx])
		if err != nil {
			return nil, err
		}
	}
	return grid, nil
}

func (g *Grid[T]) Dimensions() (int, int) {
	rows, cols := 0, 0
	if g == nil {
		return rows, cols
	}

	rows = len(g.store)
	if rows == 0 {
		return rows, cols
	}

	cols = len(g.store[0])

	return rows, cols
}

func (g *Grid[T]) InBound(c Coordinate) bool {
	return c.X >= 0 && c.X < len(g.store) && c.Y >= 0 && c.Y < len(g.store[0])
}

func (g *Grid[T]) ValueAt(c Coordinate) T {
	return g.store[c.X][c.Y]
}

func (g *Grid[T]) SetValueAt(c Coordinate, v T) {
	if !g.InBound(c) {
		panic("cannot set value in nil grid")
	}
	g.store[c.X][c.Y] = v
}

func (g *Grid[T]) IncrementValueAt(c Coordinate, v T) {
	if !g.InBound(c) {
		panic("cannot set value in nil grid")
	}
	g.store[c.X][c.Y] += v
}

func (g *Grid[T]) Find(val T) *Coordinate {
	if g == nil {
		panic("cannot find in nil grid")
	}
	for i := 0; i < len(g.store); i++ {
		for j := 0; j < len(g.store[i]); j++ {
			if g.store[i][j] == val {
				return NewCoordinate(i, j).Ptr()
			}
		}
	}
	return nil
}

func (g *Grid[T]) Print() {
	if g == nil || g.store == nil {
		fmt.Println("<nil>")
	}
	util.PrintMatrix(g.store)
}

func (g *Grid[T]) Clone() *Grid[T] {
	return &Grid[T]{
		store: util.CopyMatrix(g.store),
	}
}
