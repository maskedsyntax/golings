// channel_buffering_01: A buffered queue (reference solution)
package queue

// Queue returns a channel that already holds every message, in order.
func Queue(msgs ...string) chan string {
	ch := make(chan string, len(msgs)) // room for every message
	for _, m := range msgs {
		ch <- m
	}
	return ch
}
