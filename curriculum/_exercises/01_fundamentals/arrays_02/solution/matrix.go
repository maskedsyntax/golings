// arrays_02: Matrices as arrays (reference solution)
package matrix

// Transpose swaps rows and columns.
func Transpose(m [3][3]int) [3][3]int {
	var t [3][3]int
	for i := range m {
		for j := range m[i] {
			t[j][i] = m[i][j]
		}
	}
	return t
}

// Equal reports whether a and b hold the same values.
func Equal(a, b [3][3]int) bool {
	return a == b
}
