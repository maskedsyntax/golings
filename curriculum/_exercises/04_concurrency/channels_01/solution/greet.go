// channels_01: Sending on a channel (reference solution)
package greet

// Greet returns a channel that delivers "Hello, <name>!".
func Greet(name string) <-chan string {
	ch := make(chan string)
	go func() {
		ch <- "Hello, " + name + "!" // blocks until the caller receives
	}()
	return ch
}
