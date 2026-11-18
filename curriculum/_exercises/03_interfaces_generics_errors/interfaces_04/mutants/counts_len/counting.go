// Mutant: counts len(p) even when the wrapped writer wrote less.
package counting

import "io"

type CountingWriter struct {
	W io.Writer
	N int64
}

func (c *CountingWriter) Write(p []byte) (int, error) {
	c.N += int64(len(p))
	return c.W.Write(p)
}
