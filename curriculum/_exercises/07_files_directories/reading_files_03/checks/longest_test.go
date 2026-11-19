package longest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func write(t *testing.T, content string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "f.txt")
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestLongestLineShort(t *testing.T) {
	if n, err := LongestLine(write(t, "a\nabc\nab\n")); err != nil || n != 3 {
		t.Fatalf("LongestLine = %d, %v; want 3, nil", n, err)
	}
}

func TestLongestLineLong(t *testing.T) {
	p := write(t, "short\n"+strings.Repeat("x", 200_000)+"\nend\n")
	if n, err := LongestLine(p); err != nil || n != 200_000 {
		t.Fatalf("LongestLine = %d, %v; want 200000, nil", n, err)
	}
}

func TestLongestLineTooLong(t *testing.T) {
	p := write(t, "short\n"+strings.Repeat("x", 2<<20)+"\n")
	if n, err := LongestLine(p); err == nil {
		t.Fatalf("LongestLine = %d, nil; a 2 MiB line should be an error", n)
	}
}

func TestLongestLineMissing(t *testing.T) {
	if _, err := LongestLine(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatal("LongestLine on a missing file returned no error")
	}
}
