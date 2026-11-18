// select_01: Whichever is ready (reference solution)
package receive

// Receive returns the next n messages from a and b, in the order they arrive.
func Receive(a, b <-chan string, n int) []string {
	out := make([]string, 0, n)
	for range n {
		select {
		case m := <-a:
			out = append(out, m)
		case m := <-b:
			out = append(out, m)
		}
	}
	return out
}
