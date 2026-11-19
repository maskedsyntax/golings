package write

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteLines(t *testing.T) {
	p := filepath.Join(t.TempDir(), "out.txt")
	if err := WriteLines(p, []string{"alpha", "beta"}); err != nil {
		t.Fatalf("WriteLines returned an error: %v", err)
	}
	data, _ := os.ReadFile(p)
	if got := string(data); got != "alpha\nbeta\n" {
		t.Fatalf("file contains %q, want %q", got, "alpha\nbeta\n")
	}
}

func TestWriteLinesMany(t *testing.T) {
	p := filepath.Join(t.TempDir(), "out.txt")
	var lines []string
	var want strings.Builder
	for i := range 5000 {
		l := fmt.Sprintf("line %d", i)
		lines = append(lines, l)
		want.WriteString(l + "\n")
	}
	if err := WriteLines(p, lines); err != nil {
		t.Fatalf("WriteLines returned an error: %v", err)
	}
	data, _ := os.ReadFile(p)
	if string(data) != want.String() {
		t.Fatalf("file has %d bytes, want %d", len(data), want.Len())
	}
}

func TestWriteLinesBadDir(t *testing.T) {
	p := filepath.Join(t.TempDir(), "no", "such", "dir", "out.txt")
	if err := WriteLines(p, []string{"x"}); err == nil {
		t.Fatal("WriteLines into a missing directory returned no error")
	}
}
