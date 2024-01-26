package list

import "fmt"

type queue struct {
	store []any
}

func NewQueue() *queue {
	return &queue{make([]any, 0)}
}

func (s *queue) Push(in any) {
	s.store = append(s.store, in)
}

func (s *queue) Pop() any {
	n := len(s.store)
	if n == 0 {
		panic("nil queue")
	}
	v := s.store[0]
	s.store = s.store[1:]
	return v
}

func (s *queue) Peek() any {
	n := len(s.store)
	if n == 0 {
		panic("nil queue")
	}
	return s.store[0]
}

func (s *queue) Len() int {
	return len(s.store)
}

func (s *queue) Print() {
	for _, v := range s.store {
		fmt.Printf("%v \n", v)
	}
}
