// non_blocking_channel_operations_02: Slow subscribers
//
// One stuck subscriber blocks Publish for everyone. Fix it. See README.md.
//
// Reference: https://gobyexample.com/non-blocking-channel-operations
package hub

// Hub fans values out to subscribers.
type Hub struct {
	subs    []chan int
	Dropped int // values dropped because a subscriber was full
}

// Subscribe returns a new subscription with room for buffer pending values.
func (h *Hub) Subscribe(buffer int) <-chan int {
	ch := make(chan int, buffer)
	h.subs = append(h.subs, ch)
	return ch
}

// Publish delivers v to every subscriber that has room, without blocking.
func (h *Hub) Publish(v int) {
	for _, s := range h.subs {
		s <- v
	}
}
