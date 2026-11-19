package peek

import (
	"path/filepath"
	"testing"
)

var alphabet = filepath.Join("data", "alphabet.txt")

func TestPeek(t *testing.T) {
	for _, tt := range []struct {
		offset int64
		n      int
		want   string
	}{
		{0, 3, "abc"},
		{10, 5, "klmno"},
		{23, 10, "xyz\n"},
		{26, 1, "\n"},
		{27, 5, ""},
		{100, 5, ""},
		{5, 0, ""},
	} {
		got, err := Peek(alphabet, tt.offset, tt.n)
		if err != nil || got != tt.want {
			t.Errorf("Peek(alphabet, %d, %d) = %q, %v; want %q, nil", tt.offset, tt.n, got, err, tt.want)
		}
	}
}

func TestPeekErrors(t *testing.T) {
	if _, err := Peek(filepath.Join(t.TempDir(), "missing"), 0, 1); err == nil {
		t.Error("Peek on a missing file returned no error")
	}
	if _, err := Peek(alphabet, -1, 3); err == nil {
		t.Error("Peek with a negative offset returned no error")
	}
}
