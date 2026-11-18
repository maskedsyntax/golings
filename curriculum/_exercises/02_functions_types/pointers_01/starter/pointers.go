// pointers_01: Changing the caller's variable
//
// The tests don't compile against this file, and Swap has no effect. Fix both.
//
//	x := 4; Double(&x)       // x == 8
//	a, b := 1, 2; Swap(&a, &b) // a == 2, b == 1
//
// Reference: https://gobyexample.com/pointers
package pointers

// Double doubles the integer the caller passes in.
func Double(n int) {
	n *= 2
}

// Swap exchanges the values of *a and *b.
func Swap(a, b *int) {
	a, b = b, a
}
