package list

type heap struct {
	store     []any
	IsGreater func(a, b any) bool
}

type Heap interface {
	Push(v any)
	Pop() any
	Peek() any
	Len() int
	IsEmpty() bool
	GetStore() []any
}

/*
NewHeap retuns a new heap.
IsGreater method helps manage min/max heap
*/
func NewHeap(isGreater func(a, b any) bool) Heap {
	return &heap{
		store:     make([]any, 0),
		IsGreater: isGreater,
	}
}

func (h *heap) Push(v any) {
	h.store = append(h.store, v)
	idx := len(h.store) - 1

	for idx > 0 && h.IsGreater(h.valueAt(idx), h.valueAt(h.parent(idx))) {
		h.store[idx], h.store[h.parent(idx)] = h.store[h.parent(idx)], h.store[idx]
		idx = h.parent(idx)
	}
}

func (h *heap) Pop() any {
	if h.IsEmpty() {
		panic("empty heap")
	}

	v := h.store[0]
	h.store[0] = h.store[h.Len()-1]
	h.store = h.store[:h.Len()-1]
	h.heapify(0)

	return v
}

func (h *heap) Peek() any {
	if h.IsEmpty() {
		panic("empty heap")
	}
	return h.store[0]
}

func (h *heap) Len() int {
	return len(h.store)
}

func (h *heap) IsEmpty() bool {
	return len(h.store) == 0
}

func (h *heap) GetStore() []any {
	return h.store
}

func (h *heap) heapify(idx int) {
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

func (h *heap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *heap) leftChild(idx int) int {
	return 2*idx + 1
}

func (h *heap) rightChild(idx int) int {
	return 2*idx + 2
}

func (h *heap) valueAt(idx int) any {
	return h.store[idx]
}
