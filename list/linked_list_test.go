package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	compare := func(a, b any) bool {
		av := a.(int)
		bv := b.(int)

		return av == bv
	}

	t.Run("append to empty list", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)
		assert.Equal(t, 1, ll.Length())
		assert.NotNil(t, ll.First)
		assert.NotNil(t, ll.Last)
	})

	t.Run("append to non empty list", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)
		ll.Append(11)
		assert.Equal(t, 2, ll.Length())
	})

	t.Run("prepend to empty list", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Prepend(1)
		assert.Equal(t, 1, ll.Length())
		assert.NotNil(t, ll.First)
		assert.NotNil(t, ll.Last)
	})

	t.Run("prepend to non empty list", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)
		ll.Prepend(11)
		assert.Equal(t, 2, ll.Length())
	})

	t.Run("Delete last element", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)

		ll.Delete(3, compare)
		assert.Equal(t, 2, ll.Length())
		assert.NotNil(t, ll.First)
		assert.NotNil(t, ll.Last)
	})

	t.Run("Delete first element", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)

		ll.Delete(1, compare)
		assert.Equal(t, 2, ll.Length())
		assert.NotNil(t, ll.First)
		assert.NotNil(t, ll.Last)
	})

	t.Run("Delete only element", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)

		ll.Delete(1, compare)
		assert.Equal(t, 0, ll.Length())
		assert.Nil(t, ll.First)
		assert.Nil(t, ll.Last)
	})

}
