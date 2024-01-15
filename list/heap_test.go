package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxHeapTestCompare(a, b any) bool {
	aInt := a.(int)
	bInt := b.(int)
	return aInt > bInt
}

func minHeapTestCompare(a, b any) bool {
	aInt := a.(int)
	bInt := b.(int)
	return aInt < bInt
}

func TestHeapPush(t *testing.T) {

	t.Run("insert into empty heap", func(t *testing.T) {
		h := NewHeap(maxHeapTestCompare)
		h.Push(1)
		assert.Equal(t, []any{1}, h.GetStore())
	})

	t.Run("insert into non empty max heap", func(t *testing.T) {
		h := NewHeap(maxHeapTestCompare)
		h.Push(1)
		h.Push(2)
		h.Push(11)
		h.Push(12)
		assert.Equal(t, []any{12, 11, 2, 1}, h.GetStore())
	})

	t.Run("insert into non empty min heap", func(t *testing.T) {
		h := NewHeap(minHeapTestCompare)
		h.Push(1)
		h.Push(2)
		h.Push(11)
		h.Push(0)
		assert.Equal(t, []any{0, 1, 11, 2}, h.GetStore())
	})
}

func TestHeapPop(t *testing.T) {

	t.Run("pop from empty heap", func(t *testing.T) {
		h := NewHeap(maxHeapTestCompare)
		assert.Panics(t, func() {
			h.Pop()
		})
	})

	t.Run("pop from non empty max heap", func(t *testing.T) {
		h := NewHeap(maxHeapTestCompare)
		h.Push(1)
		h.Push(2)
		h.Push(11)
		h.Push(0)

		res := h.Pop()
		assert.Equal(t, 11, res.(int))
		assert.Equal(t, []any{2, 1, 0}, h.GetStore())
	})

	t.Run("pop from non empty min heap", func(t *testing.T) {
		h := NewHeap(minHeapTestCompare)
		h.Push(1)
		h.Push(2)
		h.Push(11)
		h.Push(0)

		res := h.Pop()
		assert.Equal(t, 0, res.(int))
		assert.Equal(t, []any{1, 2, 11}, h.GetStore())
	})
}
