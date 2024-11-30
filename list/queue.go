package list

import "fmt"

type Queue[T any] interface {
	Push(in T)
	Pop() T
	Peek() T
	Len() int
	Print()
}

type queue[T any] struct {
	store []T
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{make([]T, 0)}
}

func (s *queue[T]) Push(in T) {
	s.store = append(s.store, in)
}

func (s *queue[T]) Pop() T {
	n := len(s.store)
	if n == 0 {
		panic("nil queue")
	}
	v := s.store[0]
	s.store = s.store[1:]
	return v
}

func (s *queue[T]) Peek() T {
	n := len(s.store)
	if n == 0 {
		panic("nil queue")
	}
	return s.store[0]
}

func (s *queue[T]) Len() int {
	return len(s.store)
}

func (s *queue[T]) Print() {
	for _, v := range s.store {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
