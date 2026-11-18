package watcher

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var fast = Options{Interval: 10 * time.Millisecond, Debounce: 40 * time.Millisecond}

func expectEvent(t *testing.T, ch <-chan struct{}, want bool) {
	t.Helper()
	select {
	case <-ch:
		if !want {
			t.Fatal("unexpected change notification")
		}
	case <-time.After(400 * time.Millisecond):
		if want {
			t.Fatal("no change notification")
		}
	}
}

func TestWatchReportsContentChanges(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "a.go")
	os.WriteFile(f, []byte("one"), 0o644)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := Watch(ctx, []string{f, filepath.Join(dir, "missing.go")}, fast)
	time.Sleep(30 * time.Millisecond)

	os.WriteFile(f, []byte("two"), 0o644)
	expectEvent(t, ch, true)

	// Rewriting identical content (an editor "touch") is not a change.
	future := time.Now().Add(time.Hour)
	os.Chtimes(f, future, future)
	os.WriteFile(f, []byte("two"), 0o644)
	expectEvent(t, ch, false)

	// Deleting and creating files are changes.
	os.Remove(f)
	expectEvent(t, ch, true)
	os.WriteFile(filepath.Join(dir, "missing.go"), []byte("x"), 0o644)
	expectEvent(t, ch, true)
}

func TestWatchDebouncesBursts(t *testing.T) {
	f := filepath.Join(t.TempDir(), "a.go")
	os.WriteFile(f, []byte("0"), 0o644)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := Watch(ctx, []string{f}, Options{Interval: 10 * time.Millisecond, Debounce: 150 * time.Millisecond})
	time.Sleep(30 * time.Millisecond)
	for i := range 5 {
		os.WriteFile(f, []byte{byte('a' + i)}, 0o644)
		time.Sleep(20 * time.Millisecond)
	}
	expectEvent(t, ch, true)
	expectEvent(t, ch, false) // the burst produced exactly one notification
}

func TestWatchClosesOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := Watch(ctx, nil, fast)
	cancel()
	select {
	case _, ok := <-ch:
		if ok {
			t.Fatal("got a notification instead of close")
		}
	case <-time.After(time.Second):
		t.Fatal("channel not closed after cancel")
	}
}
