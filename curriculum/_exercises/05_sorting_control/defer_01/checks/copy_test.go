package copier

import (
	"errors"
	"slices"
	"testing"
)

type fake struct {
	data    []string
	written []string
	failAt  int
	closed  int
}

func (f *fake) Read() (string, bool) {
	if len(f.data) == 0 {
		return "", false
	}
	s := f.data[0]
	f.data = f.data[1:]
	return s, true
}

func (f *fake) Write(s string) error {
	if f.failAt > 0 && len(f.written)+1 == f.failAt {
		return errors.New("disk full")
	}
	f.written = append(f.written, s)
	return nil
}

func (f *fake) Close() { f.closed++ }

func TestCopy(t *testing.T) {
	src := &fake{data: []string{"a", "b", "c"}}
	dst := &fake{}
	if err := Copy(func() Resource { return src }, func() Resource { return dst }); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(dst.written, []string{"a", "b", "c"}) {
		t.Fatalf("wrote %q", dst.written)
	}
	if src.closed != 1 || dst.closed != 1 {
		t.Fatalf("closed source %d time(s) and destination %d time(s), want once each", src.closed, dst.closed)
	}
}

func TestCopyClosesOnError(t *testing.T) {
	src := &fake{data: []string{"a", "b", "c"}}
	dst := &fake{failAt: 2}
	if err := Copy(func() Resource { return src }, func() Resource { return dst }); err == nil {
		t.Fatal("Copy should return the write error")
	}
	if src.closed != 1 || dst.closed != 1 {
		t.Fatalf("after a failed write: closed source %d time(s) and destination %d time(s), want once each", src.closed, dst.closed)
	}
}
