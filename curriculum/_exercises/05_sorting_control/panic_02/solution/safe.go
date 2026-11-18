// panic_02: Panics you didn't ask for (reference solution)
package safe

// Last returns the last element of xs, or false if xs is empty.
func Last(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	return xs[len(xs)-1], true
}

// Registry counts registrations by name. The zero value is ready to use.
type Registry struct {
	counts map[string]int
}

// Add registers name once more.
func (r *Registry) Add(name string) {
	if r.counts == nil {
		r.counts = map[string]int{} // allocate on first write
	}
	r.counts[name]++
}

// Count returns how many times name was registered.
func (r *Registry) Count(name string) int {
	return r.counts[name] // reading a nil map is fine: it returns 0
}
