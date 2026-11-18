// panic_02: Panics you didn't ask for
//
// Both functions panic on reasonable input. Fix them. See README.md.
//
// Reference: https://gobyexample.com/panic
package safe

// Last returns the last element of xs, or false if xs is empty.
func Last(xs []int) (int, bool) {
	return xs[len(xs)-1], true
}

// Registry counts registrations by name. The zero value is ready to use.
type Registry struct {
	counts map[string]int
}

// Add registers name once more.
func (r *Registry) Add(name string) {
	r.counts[name]++
}

// Count returns how many times name was registered.
func (r *Registry) Count(name string) int {
	return r.counts[name]
}
