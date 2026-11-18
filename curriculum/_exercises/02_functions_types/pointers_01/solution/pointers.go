// pointers_01: Changing the caller's variable (reference solution)
package pointers

// Double doubles the integer the caller passes in.
func Double(p *int) {
	*p *= 2
}

// Swap exchanges the values of *a and *b.
func Swap(a, b *int) {
	*a, *b = *b, *a
}
