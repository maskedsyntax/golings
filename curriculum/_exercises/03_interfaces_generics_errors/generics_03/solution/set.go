// generics_03: A generic set (reference solution)
package set

// Set is a set of comparable values. The zero value is an empty set.
type Set[T comparable] struct {
	m map[T]struct{}
}

// Add adds v to the set.
func (s *Set[T]) Add(v T) {
	if s.m == nil {
		s.m = map[T]struct{}{}
	}
	s.m[v] = struct{}{}
}

// Has reports whether v is in the set.
func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

// Len returns the number of elements.
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Union returns a new set holding the elements of a and b.
func Union[T comparable](a, b *Set[T]) *Set[T] {
	u := &Set[T]{}
	for _, s := range []*Set[T]{a, b} {
		for v := range s.m {
			u.Add(v)
		}
	}
	return u
}
