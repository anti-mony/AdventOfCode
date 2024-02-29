package list

type Stack[T any] interface {
	Push(in T)
	Pop() T
	Peek() T
	Len() int
}

type stack[T any] struct {
	store []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{make([]T, 0)}
}

func (s *stack[T]) Push(in T) {
	s.store = append(s.store, in)
}

func (s *stack[T]) Pop() T {
	n := len(s.store)
	if n == 0 {
		panic("nil Stack")
	}
	v := s.store[n-1]
	s.store = s.store[:n-1]

	return v
}

func (s *stack[T]) Peek() T {
	n := len(s.store)
	if n == 0 {
		panic("nil Stack")
	}
	return s.store[n-1]
}

func (s *stack[T]) Len() int {
	return len(s.store)
}
