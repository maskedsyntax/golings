package atomic

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func read(t *testing.T, p string) string {
	t.Helper()
	data, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func onlyFile(t *testing.T, dir, name string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != name {
		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		t.Fatalf("directory contains %q; want only %q (no leftover temporary files)", names, name)
	}
}

func TestWriteNew(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "config.json")
	if err := WriteFileAtomic(p, strings.NewReader(`{"v":1}`)); err != nil {
		t.Fatalf("WriteFileAtomic returned an error: %v", err)
	}
	if got := read(t, p); got != `{"v":1}` {
		t.Fatalf("file contains %q", got)
	}
	onlyFile(t, dir, "config.json")
}

func TestOverwrite(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "config.json")
	os.WriteFile(p, []byte("old contents that are longer"), 0o644)
	if err := WriteFileAtomic(p, strings.NewReader("new")); err != nil {
		t.Fatalf("WriteFileAtomic returned an error: %v", err)
	}
	if got := read(t, p); got != "new" {
		t.Fatalf("file contains %q, want \"new\"", got)
	}
	onlyFile(t, dir, "config.json")
}

var errSource = errors.New("source failed")

type failAfter struct{ r io.Reader }

func (f *failAfter) Read(p []byte) (int, error) {
	n, err := f.r.Read(p)
	if err == io.EOF {
		return n, errSource
	}
	return n, err
}

func TestFailedSourceKeepsOriginal(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "config.json")
	os.WriteFile(p, []byte("old"), 0o644)
	err := WriteFileAtomic(p, &failAfter{strings.NewReader("partial new data")})
	if !errors.Is(err, errSource) {
		t.Fatalf("WriteFileAtomic error = %v, want the reader's error", err)
	}
	if got := read(t, p); got != "old" {
		t.Fatalf("after a failed write the file contains %q; it should still be \"old\"", got)
	}
	onlyFile(t, dir, "config.json")
}

func TestMissingDir(t *testing.T) {
	p := filepath.Join(t.TempDir(), "missing", "config.json")
	if err := WriteFileAtomic(p, strings.NewReader("x")); err == nil {
		t.Fatal("WriteFileAtomic into a missing directory returned no error")
	}
}
