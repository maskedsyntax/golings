// interfaces_04: Implementing io.Writer
//
// Add a Write method so *CountingWriter satisfies io.Writer. See README.md.
//
//	var buf bytes.Buffer
//	cw := &CountingWriter{W: &buf}
//	fmt.Fprintf(cw, "hello") // buf holds "hello", cw.N == 5
//
// Reference: https://gobyexample.com/interfaces
package counting

import "io"

// CountingWriter forwards writes to W and counts the bytes written.
type CountingWriter struct {
	W io.Writer
	N int64
}

// TODO: func (c *CountingWriter) Write(p []byte) (n int, err error)
