// Mutant: releases once but never reports ErrClosed.
package conn

import (
	"errors"
	"sync/atomic"
)

var ErrClosed = errors.New("connection already closed")

type Conn struct {
	closed  atomic.Bool
	release func()
}

func New(release func()) *Conn { return &Conn{release: release} }

func (c *Conn) Close() error {
	if c.closed.CompareAndSwap(false, true) {
		c.release()
	}
	return nil
}
