package list

type stack struct {
	store []any
}

func NewStack() *stack {
	return &stack{make([]any, 0)}
}

func (s *stack) Push(in any) {
	s.store = append(s.store, in)
}

func (s *stack) Pop() any {
	n := len(s.store)
	if n == 0 {
		panic("nil stack")
	}
	v := s.store[n-1]
	s.store = s.store[:n-1]

	return v
}

func (s *stack) Peek() any {
	n := len(s.store)
	if n == 0 {
		panic("nil stack")
	}
	return s.store[n-1]
}

func (s *stack) Len() int {
	return len(s.store)
}
