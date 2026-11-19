package size

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func tree(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	for name, size := range map[string]int{
		"a.txt":                  10,
		"sub/b.txt":              20,
		"sub/deeper/c.txt":       30,
		".git/objects/blob":      1000,
		"sub/.cache/tmp":         500,
		".hidden-file-is-ok.txt": 5,
	} {
		p := filepath.Join(dir, filepath.FromSlash(name))
		os.MkdirAll(filepath.Dir(p), 0o755)
		if err := os.WriteFile(p, []byte(strings.Repeat("x", size)), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestTotalSize(t *testing.T) {
	got, err := TotalSize(tree(t))
	if err != nil || got != 65 {
		t.Fatalf("TotalSize = %d, %v; want 65, nil (hidden directories are skipped, hidden files are not)", got, err)
	}
}

func TestTotalSizeDot(t *testing.T) {
	t.Chdir(tree(t))
	got, err := TotalSize(".")
	if err != nil || got != 65 {
		t.Fatalf("TotalSize(\".\") = %d, %v; want 65, nil", got, err)
	}
}

func TestTotalSizeMissing(t *testing.T) {
	if _, err := TotalSize(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatal("TotalSize(missing) returned no error")
	}
}
