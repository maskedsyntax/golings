package stream

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestHashReader(t *testing.T) {
	sum, n, err := HashReader(strings.NewReader(strings.Repeat("a", 1_000_000)))
	if err != nil {
		t.Fatalf("HashReader returned an error: %v", err)
	}
	if want := "cdc76e5c9914fb9281a1c7e284d73e67f1809a48a497200e046d39ccc7112cd0"; sum != want {
		t.Errorf("sum = %s, want %s", sum, want)
	}
	if n != 1_000_000 {
		t.Errorf("n = %d, want 1000000", n)
	}

	sum, n, err = HashReader(strings.NewReader(""))
	if err != nil || n != 0 || sum != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Errorf("HashReader(empty) = %s, %d, %v", sum, n, err)
	}
}

var errDisk = errors.New("disk on fire")

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errDisk }

func TestHashReaderError(t *testing.T) {
	_, _, err := HashReader(io.MultiReader(strings.NewReader("abc"), failingReader{}))
	if !errors.Is(err, errDisk) {
		t.Fatalf("HashReader error = %v, want the reader's error", err)
	}
}
