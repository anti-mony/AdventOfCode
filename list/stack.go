package list

type Stack[T any] struct {
	store []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (s *Stack[T]) Push(in T) {
	s.store = append(s.store, in)
}

func (s *Stack[T]) Pop() T {
	n := len(s.store)
	if n == 0 {
		panic("nil Stack")
	}
	v := s.store[n-1]
	s.store = s.store[:n-1]

	return v
}

func (s *Stack[T]) Peek() T {
	n := len(s.store)
	if n == 0 {
		panic("nil Stack")
	}
	return s.store[n-1]
}

func (s *Stack[T]) Len() int {
	return len(s.store)
}
