// Mutant: Union adds b's elements into a and returns a.
package set

type Set[T comparable] struct {
	m map[T]struct{}
}

func (s *Set[T]) Add(v T) {
	if s.m == nil {
		s.m = map[T]struct{}{}
	}
	s.m[v] = struct{}{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Len() int { return len(s.m) }

func Union[T comparable](a, b *Set[T]) *Set[T] {
	for v := range b.m {
		a.Add(v)
	}
	return a
}
