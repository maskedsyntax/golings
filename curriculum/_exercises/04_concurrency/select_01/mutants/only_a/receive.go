// Mutant: ignores b.
package receive

func Receive(a, b <-chan string, n int) []string {
	var out []string
	for range n {
		out = append(out, <-a)
	}
	return out
}
