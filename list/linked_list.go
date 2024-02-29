package list

import "fmt"

type node[T any] struct {
	Value    T
	next     *node[T]
	previous *node[T]
}

type LinkedList[T any] struct {
	First *node[T]
	Last  *node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(value T) {
	node := node[T]{Value: value, next: nil, previous: nil}

	if l.First == nil && l.Last == nil {
		l.First = &node
		l.Last = &node
		return
	}

	node.previous = l.Last
	l.Last.next = &node

	l.Last = &node
}

func (l *LinkedList[T]) Prepend(value T) {
	node := node[T]{Value: value, next: nil, previous: nil}

	if l.First == nil && l.Last == nil {
		l.First = &node
		l.Last = &node
		return
	}

	node.next = l.First
	l.First.previous = &node

	l.First = &node
}

func (l *LinkedList[T]) Print() {
	tmp := l.First

	for tmp != nil {
		fmt.Printf("%v --> ", tmp.Value)
		tmp = tmp.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) Delete(v T, compare func(a, b T) bool) {
	if l.IsEmpty() {
		return
	}
	if compare(l.First.Value, v) {
		if l.First.next == nil {
			l.First = nil
			l.Last = nil
			return
		}
		l.First = l.First.next
		return
	}

	tmp := l.First

	for tmp != nil {
		if compare(tmp.Value, v) {
			tmp.previous.next = tmp.next
			if tmp.next == nil {
				l.Last = tmp.previous
				return
			}
			tmp.next.previous = tmp.previous
			return
		}
		tmp = tmp.next
	}
}

func (l *LinkedList[T]) UpsertAppend(v T, compare func(a, b T) bool) {
	if l.First == nil {
		l.Append(v)
		return
	}

	tmp := l.First

	for tmp != nil {
		if compare(tmp.Value, v) {
			tmp.Value = v
			return
		}
		tmp = tmp.next
	}

	l.Append(v)
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.First == nil
}

func (l *LinkedList[T]) Length() int {
	length := 0
	tmp := l.First
	for tmp != nil {
		length++
		tmp = tmp.next
	}
	return length
}
