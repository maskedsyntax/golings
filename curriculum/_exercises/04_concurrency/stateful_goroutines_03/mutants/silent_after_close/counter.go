// Mutant: no longer hangs after Close, but reports success.
package counter

import "errors"

var ErrClosed = errors.New("counter closed")

type Counter struct {
	incs chan chan int
	quit chan struct{}
	done chan struct{}
}

func NewCounter() *Counter {
	c := &Counter{incs: make(chan chan int), quit: make(chan struct{}), done: make(chan struct{})}
	go func() {
		defer close(c.done)
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

func (c *Counter) Inc() (int, error) {
	resp := make(chan int)
	select {
	case c.incs <- resp:
		return <-resp, nil
	case <-c.done:
		return 0, nil
	}
}

func (c *Counter) Close() {
	close(c.quit)
	<-c.done
}
