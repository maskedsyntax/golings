// atomic_counters_03: Close exactly once (reference solution)
package conn

import (
	"errors"
	"sync/atomic"
)

// ErrClosed is returned by Close after the first call.
var ErrClosed = errors.New("connection already closed")

// Conn is a connection that must be released exactly once.
type Conn struct {
	closed  atomic.Bool
	release func()
}

// New returns a Conn that calls release when closed.
func New(release func()) *Conn { return &Conn{release: release} }

// Close releases the connection the first time it is called.
func (c *Conn) Close() error {
	if !c.closed.CompareAndSwap(false, true) {
		return ErrClosed // someone else won
	}
	c.release()
	return nil
}
