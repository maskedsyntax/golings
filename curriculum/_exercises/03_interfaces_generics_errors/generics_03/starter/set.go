// generics_03: A generic set
//
// Implement the Set methods and Union. See README.md.
//
//	var s Set[string]
//	s.Add("go"); s.Add("go")
//	s.Len() == 1; s.Has("go") == true
//
// Reference: https://gobyexample.com/generics
package set

// Set is a set of comparable values. The zero value is an empty set.
type Set[T comparable] struct {
	m map[T]struct{}
}

// Add adds v to the set.
func (s *Set[T]) Add(v T) {
	// TODO
}

// Has reports whether v is in the set.
func (s *Set[T]) Has(v T) bool {
	// TODO
	return false
}

// Len returns the number of elements.
func (s *Set[T]) Len() int {
	// TODO
	return 0
}

// Union returns a new set holding the elements of a and b.
func Union[T comparable](a, b *Set[T]) *Set[T] {
	// TODO
	return &Set[T]{}
}
