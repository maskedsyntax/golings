// channels_01: Sending on a channel
//
// Greet blocks forever. Fix it. See README.md.
//
// Reference: https://gobyexample.com/channels
package greet

// Greet returns a channel that delivers "Hello, <name>!".
func Greet(name string) <-chan string {
	ch := make(chan string)
	ch <- "Hello, " + name + "!"
	return ch
}
