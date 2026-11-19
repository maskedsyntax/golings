package grep

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	for _, tt := range []struct {
		in, pattern, want string
		n                 int
	}{
		{"alpha\r\nbeta\r\ngamma alpha\n", "alpha", "1:alpha\n3:gamma alpha\n", 2},
		{"one\ntwo\nthree", "t", "2:two\n3:three\n", 2},
		{"x\ny\n", "z", "", 0},
		{"Go\ngo\n", "go", "2:go\n", 1},
	} {
		var out strings.Builder
		n, err := Grep(strings.NewReader(tt.in), &out, tt.pattern)
		if err != nil {
			t.Fatalf("Grep(%q, %q) returned an error: %v", tt.in, tt.pattern, err)
		}
		if out.String() != tt.want || n != tt.n {
			t.Errorf("Grep(%q, %q) wrote %q and returned %d; want %q and %d", tt.in, tt.pattern, out.String(), n, tt.want, tt.n)
		}
	}
}

var errRead = errors.New("read failed")

type failing struct{}

func (failing) Read([]byte) (int, error) { return 0, errRead }

func TestGrepReadError(t *testing.T) {
	_, err := Grep(io.MultiReader(strings.NewReader("a\n"), failing{}), io.Discard, "a")
	if !errors.Is(err, errRead) {
		t.Fatalf("Grep error = %v, want the reader's error", err)
	}
}
