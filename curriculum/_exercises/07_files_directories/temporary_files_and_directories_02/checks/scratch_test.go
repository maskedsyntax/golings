package scratch

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fill(t *testing.T, dir string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(dir, "a", "b"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "a", "b", "f.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func gone(t *testing.T, dir string) {
	t.Helper()
	if dir == "" {
		t.Fatal("fn was never called")
	}
	if _, err := os.Stat(dir); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("%s still exists after Scratch returned", dir)
	}
}

func TestScratch(t *testing.T) {
	var seen string
	err := Scratch(func(dir string) error {
		seen = dir
		if fi, err := os.Stat(dir); err != nil || !fi.IsDir() {
			t.Errorf("%s is not a directory while fn runs", dir)
		}
		if !strings.HasPrefix(filepath.Base(dir), "scratch-") {
			t.Errorf("directory name %q should start with scratch-", filepath.Base(dir))
		}
		fill(t, dir)
		return nil
	})
	if err != nil {
		t.Fatalf("Scratch returned an error: %v", err)
	}
	gone(t, seen)
}

func TestScratchError(t *testing.T) {
	boom := errors.New("boom")
	var seen string
	err := Scratch(func(dir string) error { seen = dir; fill(t, dir); return boom })
	if !errors.Is(err, boom) {
		t.Fatalf("Scratch error = %v, want fn's error", err)
	}
	gone(t, seen)
}

func TestScratchPanic(t *testing.T) {
	var seen string
	func() {
		defer func() { recover() }()
		Scratch(func(dir string) error { seen = dir; fill(t, dir); panic("oops") })
	}()
	gone(t, seen)
}
