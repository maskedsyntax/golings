// multiple_return_values_01: Receiving two results
//
// This file does not compile. Fix it.
//
//	DivMod(7, 2) == 3, 1
//	Quotient(7, 2) == 3
//
// Reference: https://gobyexample.com/multiple-return-values
package divmod

// DivMod returns the quotient and remainder of a / b.
func DivMod(a, b int) int, int {
	return a / b, a % b
}

// Quotient returns a / b, computed with DivMod.
func Quotient(a, b int) int {
	q := DivMod(a, b)
	return q
}
