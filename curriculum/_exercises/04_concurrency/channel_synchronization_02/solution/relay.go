// channel_synchronization_02: Relay race (reference solution)
package relay

// Relay runs each step in its own goroutine, strictly in order, and returns
// when the last one has finished.
func Relay(steps ...func()) {
	prev := make(chan struct{})
	close(prev) // the first runner may start immediately
	for _, step := range steps {
		done := make(chan struct{})
		go func(wait <-chan struct{}) {
			<-wait // the baton from the previous runner
			step()
			close(done)
		}(prev)
		prev = done
	}
	<-prev // the last runner has finished
}
