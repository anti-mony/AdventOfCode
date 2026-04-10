package list

type Set[T comparable] interface {
	Add(in T)
	Remove(in T)
	Size() int
	Contains(in T) bool
}

type _token struct{}

type set[T comparable] struct {
	store map[T]_token
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{
		store: make(map[T]_token),
	}
}

func (s *set[T]) Add(in T) {
	s.store[in] = _token{}
}

func (s *set[T]) Remove(in T) {
	delete(s.store, in)
}

func (s *set[T]) Contains(in T) bool {
	_, ok := s.store[in]
	return ok
}

func (s *set[T]) Size() int {
	return len(s.store)
}
