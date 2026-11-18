// atomic_counters_03: Close exactly once
//
// Close is not safe for concurrent use. Fix it with sync/atomic.
// See README.md.
//
// Reference: https://gobyexample.com/atomic-counters
package conn

import "errors"

// ErrClosed is returned by Close after the first call.
var ErrClosed = errors.New("connection already closed")

// Conn is a connection that must be released exactly once.
type Conn struct {
	closed  bool
	release func()
}

// New returns a Conn that calls release when closed.
func New(release func()) *Conn { return &Conn{release: release} }

// Close releases the connection the first time it is called.
func (c *Conn) Close() error {
	if c.closed {
		return ErrClosed
	}
	c.closed = true
	c.release()
	return nil
}
