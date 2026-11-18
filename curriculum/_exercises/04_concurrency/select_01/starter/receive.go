// select_01: Whichever is ready
//
// Receive gets stuck when the channels aren't taking turns. Use select.
// See README.md.
//
// Reference: https://gobyexample.com/select
package receive

// Receive returns the next n messages from a and b, in the order they arrive.
func Receive(a, b <-chan string, n int) []string {
	var out []string
	for i := range n {
		if i%2 == 0 {
			out = append(out, <-a)
		} else {
			out = append(out, <-b)
		}
	}
	return out
}
