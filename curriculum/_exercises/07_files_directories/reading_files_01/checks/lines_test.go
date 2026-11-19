package lines

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestCountLinesFixture(t *testing.T) {
	n, err := CountLines(filepath.Join("data", "poem.txt"))
	if err != nil || n != 4 {
		t.Fatalf("CountLines(data/poem.txt) = %d, %v; want 4, nil", n, err)
	}
}

func TestCountLines(t *testing.T) {
	dir := t.TempDir()
	for _, tt := range []struct {
		content string
		want    int
	}{
		{"a\nb", 2},
		{"a\nb\n", 2},
		{"\n\n", 2},
		{"", 0},
	} {
		p := filepath.Join(dir, "f.txt")
		if err := os.WriteFile(p, []byte(tt.content), 0o644); err != nil {
			t.Fatal(err)
		}
		if n, err := CountLines(p); err != nil || n != tt.want {
			t.Errorf("CountLines(%q) = %d, %v; want %d, nil", tt.content, n, err, tt.want)
		}
	}
}

func TestCountLinesMissing(t *testing.T) {
	_, err := CountLines(filepath.Join(t.TempDir(), "missing.txt"))
	if !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("CountLines(missing file) error = %v, want one matching fs.ErrNotExist", err)
	}
}
