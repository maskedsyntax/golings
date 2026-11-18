// channel_synchronization_01: A done signal (reference solution)
package start

// Start runs task in a new goroutine and returns a channel that is closed
// when task has finished.
func Start(task func()) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done) // runs after task returns (even if it panics)
		task()
	}()
	return done
}
