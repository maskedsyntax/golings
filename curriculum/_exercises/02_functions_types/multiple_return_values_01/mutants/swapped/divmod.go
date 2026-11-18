// Mutant: Quotient keeps the wrong one of the two results.
package divmod

func DivMod(a, b int) (int, int) { return a / b, a % b }

func Quotient(a, b int) int {
	_, r := DivMod(a, b)
	return r
}
