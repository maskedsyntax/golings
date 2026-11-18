// channel_synchronization_01: A done signal
//
// The done channel is closed too early. Fix it. See README.md.
//
// Reference: https://gobyexample.com/channel-synchronization
package start

// Start runs task in a new goroutine and returns a channel that is closed
// when task has finished.
func Start(task func()) <-chan struct{} {
	done := make(chan struct{})
	go task()
	close(done)
	return done
}
