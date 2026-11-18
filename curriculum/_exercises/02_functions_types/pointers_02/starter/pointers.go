// pointers_02: nil pointers and identity
//
// Implement ValueOr, Increment and Largest. See README.md.
//
//	ValueOr(nil, 7) == 7
//	a, b := 3, 9; Largest([]*int{&a, nil, &b}) == &b
//
// Reference: https://gobyexample.com/pointers
package pointers

// ValueOr returns *p, or def if p is nil.
func ValueOr(p *int, def int) int {
	// TODO
	return 0
}

// Increment adds one to *p. It does nothing if p is nil.
func Increment(p *int) {
	// TODO
}

// Largest returns the element of ptrs that points to the largest value,
// skipping nil entries. It returns nil if there is none.
func Largest(ptrs []*int) *int {
	// TODO
	return nil
}
