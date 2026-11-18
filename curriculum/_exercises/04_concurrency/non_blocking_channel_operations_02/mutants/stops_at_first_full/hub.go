// Mutant: gives up on all remaining subscribers once one is full.
package hub

type Hub struct {
	subs    []chan int
	Dropped int
}

func (h *Hub) Subscribe(buffer int) <-chan int {
	ch := make(chan int, buffer)
	h.subs = append(h.subs, ch)
	return ch
}

func (h *Hub) Publish(v int) {
	for _, s := range h.subs {
		select {
		case s <- v:
		default:
			h.Dropped++
			return
		}
	}
}
