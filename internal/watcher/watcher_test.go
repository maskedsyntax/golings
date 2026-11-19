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
	// Waiting for an event that should come gets generous slack for slow
	// machines; it returns as soon as the event arrives.
	wait := 400 * time.Millisecond
	if want {
		wait = 2 * time.Second
	}
	select {
	case <-ch:
		if !want {
			t.Fatal("unexpected change notification")
		}
	case <-time.After(wait):
		if want {
			t.Fatal("no change notification")
		}
	}
}

// save replaces path atomically, as most editors do. os.WriteFile truncates
// first, and on a slow machine the empty intermediate file can outlast the
// short test debounce and be reported as a change of its own.
func save(t *testing.T, path, content string) {
	t.Helper()
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(tmp, path); err != nil {
		t.Fatal(err)
	}
}

func TestWatchReportsContentChanges(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "a.go")
	save(t, f, "one")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := Watch(ctx, []string{f, filepath.Join(dir, "missing.go")}, fast)
	time.Sleep(30 * time.Millisecond)

	save(t, f, "two")
	expectEvent(t, ch, true)

	// Rewriting identical content (an editor "touch") is not a change.
	future := time.Now().Add(time.Hour)
	os.Chtimes(f, future, future)
	save(t, f, "two")
	expectEvent(t, ch, false)

	// Deleting and creating files are changes.
	os.Remove(f)
	expectEvent(t, ch, true)
	save(t, filepath.Join(dir, "missing.go"), "x")
	expectEvent(t, ch, true)
}

func TestWatchDebouncesBursts(t *testing.T) {
	f := filepath.Join(t.TempDir(), "a.go")
	os.WriteFile(f, []byte("0"), 0o644)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const debounce = 300 * time.Millisecond
	ch := Watch(ctx, []string{f}, Options{Interval: 10 * time.Millisecond, Debounce: debounce})
	time.Sleep(30 * time.Millisecond)
	last := time.Now()
	var maxGap time.Duration
	for i := range 5 {
		save(t, f, string(rune('a'+i)))
		now := time.Now()
		maxGap = max(maxGap, now.Sub(last))
		last = now
		time.Sleep(20 * time.Millisecond)
	}
	if maxGap >= debounce {
		// The machine stalled long enough that this really was two bursts.
		t.Skipf("saves were %v apart (debounce %v); too slow to test burst coalescing", maxGap, debounce)
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
