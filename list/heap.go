package list

type heap[T any] struct {
	IsGreater func(a, b T) bool
	store     []T
}

type Heap[T any] interface {
	Push(v T)
	Pop() T
	Peek() T
	Len() int
	IsEmpty() bool
	GetStore() []T
}

/*
NewHeap retuns a new heap.
IsGreater method helps manage min/max heap
*/
func NewHeap[T any](isGreater func(a, b T) bool) Heap[T] {
	return &heap[T]{
		store:     make([]T, 0),
		IsGreater: isGreater,
	}
}

func (h *heap[T]) Push(v T) {
	h.store = append(h.store, v)
	idx := len(h.store) - 1

	for idx > 0 && h.IsGreater(h.valueAt(idx), h.valueAt(h.parent(idx))) {
		h.store[idx], h.store[h.parent(idx)] = h.store[h.parent(idx)], h.store[idx]
		idx = h.parent(idx)
	}
}

func (h *heap[T]) Pop() T {
	if h.IsEmpty() {
		panic("empty heap")
	}

	v := h.store[0]
	h.store[0] = h.store[h.Len()-1]
	h.store = h.store[:h.Len()-1]
	h.heapify(0)

	return v
}

func (h *heap[T]) Peek() T {
	if h.IsEmpty() {
		panic("empty heap")
	}
	return h.store[0]
}

func (h *heap[T]) Len() int {
	return len(h.store)
}

func (h *heap[T]) IsEmpty() bool {
	return len(h.store) == 0
}

func (h *heap[T]) GetStore() []T {
	return h.store
}

func (h *heap[T]) heapify(idx int) {
	biggest := idx
	leftChild := h.leftChild(idx)
	rightChild := h.rightChild(idx)

	if leftChild < h.Len() && h.IsGreater(h.valueAt(leftChild), h.valueAt(biggest)) {
		biggest = leftChild
	}
	if rightChild < h.Len() && h.IsGreater(h.valueAt(rightChild), h.valueAt(biggest)) {
		biggest = rightChild
	}

	if biggest != idx {
		h.store[idx], h.store[biggest] = h.store[biggest], h.store[idx]
		h.heapify(biggest)
	}
}

func (h *heap[T]) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *heap[T]) leftChild(idx int) int {
	return 2*idx + 1
}

func (h *heap[T]) rightChild(idx int) int {
	return 2*idx + 2
}

func (h *heap[T]) valueAt(idx int) T {
	return h.store[idx]
}
