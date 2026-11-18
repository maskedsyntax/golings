// range_over_iterators_03: Iterating a custom collection (reference solution)
package list

import (
	"iter"
	"slices"
)

type node[T any] struct {
	val  T
	next *node[T]
}

// List is a singly linked list. The zero value is an empty list.
type List[T any] struct {
	head, tail *node[T]
}

// Push appends v to the end of the list.
func (l *List[T]) Push(v T) {
	n := &node[T]{val: v}
	if l.tail == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

// Values yields the values from front to back.
func (l *List[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := l.head; n != nil; n = n.next {
			if !yield(n.val) {
				return
			}
		}
	}
}

// All yields (index, value) pairs from front to back.
func (l *List[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for n := l.head; n != nil; n = n.next {
			if !yield(i, n.val) {
				return
			}
			i++
		}
	}
}

// Backward yields (index, value) pairs from back to front.
func (l *List[T]) Backward() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		// A singly linked list can't walk backwards, so buffer the values.
		vals := slices.Collect(l.Values())
		for i := len(vals) - 1; i >= 0; i-- {
			if !yield(i, vals[i]) {
				return
			}
		}
	}
}
