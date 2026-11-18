// Mutant: Backward yields values in reverse but numbers them 0, 1, 2...
package list

import (
	"iter"
	"slices"
)

type node[T any] struct {
	val  T
	next *node[T]
}

type List[T any] struct {
	head, tail *node[T]
}

func (l *List[T]) Push(v T) {
	n := &node[T]{val: v}
	if l.tail == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

func (l *List[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := l.head; n != nil; n = n.next {
			if !yield(n.val) {
				return
			}
		}
	}
}

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

func (l *List[T]) Backward() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		vals := slices.Collect(l.Values())
		slices.Reverse(vals)
		for i, v := range vals {
			if !yield(i, v) {
				return
			}
		}
	}
}
