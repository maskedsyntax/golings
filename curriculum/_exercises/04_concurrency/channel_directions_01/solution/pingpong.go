// channel_directions_01: Send-only and receive-only (reference solution)
package pingpong

// Ping sends msg on pings.
func Ping(pings chan<- string, msg string) {
	pings <- msg
}

// Pong receives one message from pings and passes it on to pongs.
func Pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
