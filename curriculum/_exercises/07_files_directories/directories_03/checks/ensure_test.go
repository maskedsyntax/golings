package ensure

import (
	"os"
	"path/filepath"
	"testing"
)

func isDir(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.IsDir()
}

func TestEnsureDirs(t *testing.T) {
	root := t.TempDir()
	rels := []string{"logs", filepath.Join("data", "cache", "images"), "data"}
	for i := range 2 { // a second run must succeed too
		if err := EnsureDirs(root, rels); err != nil {
			t.Fatalf("run %d: EnsureDirs returned an error: %v", i+1, err)
		}
	}
	for _, rel := range rels {
		if !isDir(filepath.Join(root, rel)) {
			t.Errorf("%s was not created", rel)
		}
	}
}

func TestEnsureDirsRejectsEscapes(t *testing.T) {
	parent := t.TempDir()
	root := filepath.Join(parent, "root")
	if err := os.Mkdir(root, 0o755); err != nil {
		t.Fatal(err)
	}
	abs := filepath.Join(parent, "absolute")
	for _, rel := range []string{filepath.Join("..", "escape"), abs, ""} {
		if err := EnsureDirs(root, []string{rel}); err == nil {
			t.Errorf("EnsureDirs(root, %q) returned no error", rel)
		}
	}
	if isDir(filepath.Join(parent, "escape")) {
		t.Error("EnsureDirs created a directory outside root")
	}
}
