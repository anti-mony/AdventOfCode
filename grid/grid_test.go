package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	t.Run("new grid", func(t *testing.T) {
		g := NewGrid[int](5, 5)

		assert.NotNil(t, g)
		assert.Len(t, g.store, 5)
		assert.Len(t, g.store[0], 5)
	})

	t.Run("dimensions, nil grid", func(t *testing.T) {
		var g *Grid[int]
		rows, cols := g.Dimensions()
		assert.Equal(t, 0, rows)
		assert.Equal(t, 0, cols)
	})

	t.Run("dimensions, no rows", func(t *testing.T) {
		g := NewGrid[int](0, 5)
		rows, cols := g.Dimensions()
		assert.Equal(t, 0, rows)
		assert.Equal(t, 0, cols)
	})

	t.Run("dimensions, no cols", func(t *testing.T) {
		g := NewGrid[int](5, 0)
		rows, cols := g.Dimensions()
		assert.Equal(t, 5, rows)
		assert.Equal(t, 0, cols)
	})

	t.Run("dimensions", func(t *testing.T) {
		g := NewGrid[int](5, 10)
		rows, cols := g.Dimensions()
		assert.Equal(t, 5, rows)
		assert.Equal(t, 10, cols)
	})

	t.Run("in bounds", func(t *testing.T) {
		g := NewGrid[int](3, 3)
		assert.True(t, g.InBound(NewCoordinate(1, 2)))
		assert.False(t, g.InBound(NewCoordinate(3, 3)))
	})

	t.Run("value at", func(t *testing.T) {
		g := NewGrid[int](5, 5)
		g.store[3][3] = 7

		assert.Equal(t, 7, g.ValueAt(NewCoordinate(3, 3)))
	})
}
