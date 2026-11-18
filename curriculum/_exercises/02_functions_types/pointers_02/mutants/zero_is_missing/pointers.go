// Mutant: treats a pointer to 0 like nil, and starts Largest at 0.
package pointers

func ValueOr(p *int, def int) int {
	if p == nil || *p == 0 {
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
	bestVal := 0
	for _, p := range ptrs {
		if p != nil && *p > bestVal {
			best, bestVal = p, *p
		}
	}
	return best
}
