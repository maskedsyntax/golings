// pointers_02: nil pointers and identity (reference solution)
package pointers

// ValueOr returns *p, or def if p is nil.
func ValueOr(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}

// Increment adds one to *p. It does nothing if p is nil.
func Increment(p *int) {
	if p != nil {
		*p++
	}
}

// Largest returns the element of ptrs that points to the largest value,
// skipping nil entries. It returns nil if there is none.
func Largest(ptrs []*int) *int {
	var best *int
	for _, p := range ptrs {
		if p != nil && (best == nil || *p > *best) {
			best = p
		}
	}
	return best
}
