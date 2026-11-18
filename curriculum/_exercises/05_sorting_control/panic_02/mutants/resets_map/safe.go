// Mutant: allocates a fresh map on every Add, losing earlier counts.
package safe

func Last(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	return xs[len(xs)-1], true
}

type Registry struct {
	counts map[string]int
}

func (r *Registry) Add(name string) {
	r.counts = map[string]int{name: r.counts[name] + 1}
}

func (r *Registry) Count(name string) int {
	return r.counts[name]
}
