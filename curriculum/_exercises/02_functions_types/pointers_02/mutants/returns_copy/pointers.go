// Mutant: Largest returns a pointer to a copy of the value.
package pointers

func ValueOr(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}

func Increment(p *int) {
	if p != nil {
		*p++
	}
}

func Largest(ptrs []*int) *int {
	var best *int
	for _, p := range ptrs {
		if p != nil && (best == nil || *p > *best) {
			v := *p
			best = &v
		}
	}
	return best
}
