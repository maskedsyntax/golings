package list

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestListGoFiles(t *testing.T) {
	dir := t.TempDir()
	for _, f := range []string{"util.go", "main.go", "util_test.go", "README.md", filepath.Join("sub", "deep.go")} {
		p := filepath.Join(dir, f)
		os.MkdirAll(filepath.Dir(p), 0o755)
		if err := os.WriteFile(p, []byte("package x\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	os.Mkdir(filepath.Join(dir, "weird.go"), 0o755) // a directory with a .go name

	got, err := ListGoFiles(dir)
	if err != nil {
		t.Fatalf("ListGoFiles returned an error: %v", err)
	}
	if want := []string{"main.go", "util.go"}; !slices.Equal(got, want) {
		t.Fatalf("ListGoFiles = %q, want %q", got, want)
	}
}

func TestListGoFilesEmpty(t *testing.T) {
	got, err := ListGoFiles(t.TempDir())
	if err != nil || len(got) != 0 {
		t.Fatalf("ListGoFiles(empty dir) = %q, %v; want nothing", got, err)
	}
}

func TestListGoFilesMissing(t *testing.T) {
	if _, err := ListGoFiles(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatal("ListGoFiles(missing dir) returned no error")
	}
}
