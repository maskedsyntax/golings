package upper

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestUpper(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{"hello\nworld\n", "HELLO\nWORLD\n"},
		{"no newline", "NO NEWLINE\n"},
		{"mixed Case 123\n\nend", "MIXED CASE 123\n\nEND\n"},
		{"", ""},
	} {
		var out strings.Builder
		if err := Upper(strings.NewReader(tt.in), &out); err != nil {
			t.Fatalf("Upper(%q) returned an error: %v", tt.in, err)
		}
		if out.String() != tt.want {
			t.Errorf("Upper(%q) wrote %q, want %q", tt.in, out.String(), tt.want)
		}
	}
}

var errRead = errors.New("read failed")

type failing struct{}

func (failing) Read([]byte) (int, error) { return 0, errRead }

func TestUpperReadError(t *testing.T) {
	var out strings.Builder
	err := Upper(io.MultiReader(strings.NewReader("ok\n"), failing{}), &out)
	if !errors.Is(err, errRead) {
		t.Fatalf("Upper error = %v, want the reader's error", err)
	}
}
