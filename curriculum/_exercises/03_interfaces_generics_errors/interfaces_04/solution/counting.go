// interfaces_04: Implementing io.Writer (reference solution)
package counting

import "io"

// CountingWriter forwards writes to W and counts the bytes written.
type CountingWriter struct {
	W io.Writer
	N int64
}

// Write forwards p to the wrapped writer and counts what it wrote.
func (c *CountingWriter) Write(p []byte) (int, error) {
	n, err := c.W.Write(p)
	c.N += int64(n)
	return n, err
}
