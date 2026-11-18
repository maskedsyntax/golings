// stateful_goroutines_03: Shutting the owner down
//
// Inc hangs after Close. Fix Counter. See README.md.
//
// Reference: https://gobyexample.com/stateful-goroutines
package counter

import "errors"

// ErrClosed is returned by Inc after Close.
var ErrClosed = errors.New("counter closed")

// Counter is a counter owned by a single goroutine.
type Counter struct {
	incs chan chan int // each request carries the channel for its reply
	quit chan struct{}
	done chan struct{} // should be closed when the owner goroutine exits
}

// NewCounter starts the owner goroutine.
func NewCounter() *Counter {
	c := &Counter{incs: make(chan chan int), quit: make(chan struct{}), done: make(chan struct{})}
	go func() {
		n := 0
		for {
			select {
			case resp := <-c.incs:
				n++
				resp <- n
			case <-c.quit:
				return
			}
		}
	}()
	return c
}

// Inc increments the counter and returns the new value.
func (c *Counter) Inc() (int, error) {
	resp := make(chan int)
	c.incs <- resp
	return <-resp, nil
}

// Close stops the owner goroutine and waits for it to exit.
func (c *Counter) Close() {
	close(c.quit)
}
