// Mutant: swapping every (i, j) pair in place swaps each pair twice.
package matrix

func Transpose(m [3][3]int) [3][3]int {
	for i := range m {
		for j := range m[i] {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	return m
}

func Equal(a, b [3][3]int) bool { return a == b }
