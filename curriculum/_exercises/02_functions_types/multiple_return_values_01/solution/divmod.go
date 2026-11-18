// multiple_return_values_01: Receiving two results (reference solution)
package divmod

// DivMod returns the quotient and remainder of a / b.
func DivMod(a, b int) (int, int) {
	return a / b, a % b
}

// Quotient returns a / b, computed with DivMod.
func Quotient(a, b int) int {
	q, _ := DivMod(a, b)
	return q
}
