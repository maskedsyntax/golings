// channel_buffering_01: A buffered queue
//
// Queue blocks. Fix it without adding a goroutine. See README.md.
//
// Reference: https://gobyexample.com/channel-buffering
package queue

// Queue returns a channel that already holds every message, in order.
func Queue(msgs ...string) chan string {
	ch := make(chan string)
	for _, m := range msgs {
		ch <- m
	}
	return ch
}
