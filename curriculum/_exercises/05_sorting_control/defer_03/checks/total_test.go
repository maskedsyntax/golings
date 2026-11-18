package total

import (
	"errors"
	"io"
	"strings"
	"testing"
)

type tracker struct{ open, peak, closes int }

type file struct {
	io.Reader
	t      *tracker
	closed bool
}

func (f *file) Close() error {
	if !f.closed {
		f.closed = true
		f.t.open--
		f.t.closes++
	}
	return nil
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("read error") }

func (tr *tracker) opener(files map[string]string) Opener {
	return func(name string) (io.ReadCloser, error) {
		content, ok := files[name]
		if !ok {
			return nil, errors.New("no such file: " + name)
		}
		tr.open++
		tr.peak = max(tr.peak, tr.open)
		var r io.Reader = strings.NewReader(content)
		if content == "<fail>" {
			r = failingReader{}
		}
		return &file{Reader: r, t: tr}, nil
	}
}

func TestTotalSize(t *testing.T) {
	tr := &tracker{}
	files := map[string]string{"a": "hello", "b": "go", "c": "gopher!"}
	n, err := TotalSize(tr.opener(files), []string{"a", "b", "c"})
	if err != nil || n != 14 {
		t.Fatalf("TotalSize = %d, %v; want 14, nil", n, err)
	}
	if tr.peak != 1 {
		t.Fatalf("%d files were open at the same time, want at most 1", tr.peak)
	}
	if tr.open != 0 {
		t.Fatalf("%d file(s) left open", tr.open)
	}
}

func TestTotalSizeClosesOnReadError(t *testing.T) {
	tr := &tracker{}
	files := map[string]string{"a": "x", "bad": "<fail>"}
	if _, err := TotalSize(tr.opener(files), []string{"a", "bad"}); err == nil {
		t.Fatal("TotalSize should return the read error")
	}
	if tr.open != 0 {
		t.Fatalf("%d file(s) left open after a read error", tr.open)
	}
}
