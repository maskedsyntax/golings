package counting

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

var _ io.Writer = (*CountingWriter)(nil)

func TestCountsAndForwards(t *testing.T) {
	var buf bytes.Buffer
	cw := &CountingWriter{W: &buf}
	fmt.Fprintf(cw, "hello, %s!", "world")
	io.Copy(cw, strings.NewReader(" and more"))
	if got := buf.String(); got != "hello, world! and more" {
		t.Errorf("wrapped writer received %q", got)
	}
	if cw.N != int64(len("hello, world! and more")) {
		t.Errorf("N = %d, want %d", cw.N, len("hello, world! and more"))
	}
}

// shortWriter accepts at most limit bytes in total, then fails.
type shortWriter struct{ limit int }

var errFull = errors.New("disk full")

func (w *shortWriter) Write(p []byte) (int, error) {
	if len(p) <= w.limit {
		w.limit -= len(p)
		return len(p), nil
	}
	n := w.limit
	w.limit = 0
	return n, errFull
}

func TestShortWrite(t *testing.T) {
	cw := &CountingWriter{W: &shortWriter{limit: 3}}
	n, err := cw.Write([]byte("hello"))
	if n != 3 || !errors.Is(err, errFull) {
		t.Errorf("Write returned %d, %v, want 3, %v", n, err, errFull)
	}
	if cw.N != 3 {
		t.Errorf("N = %d after a short write of 3 bytes, want 3", cw.N)
	}
}
