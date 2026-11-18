// Mutant: Double is fixed, but Swap still only swaps its local pointers.
package pointers

func Double(p *int) { *p *= 2 }

func Swap(a, b *int) {
	a, b = b, a
}
