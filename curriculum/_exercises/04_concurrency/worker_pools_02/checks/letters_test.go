package letters

import (
	"maps"
	"sync/atomic"
	"testing"
	"time"
	"unicode/utf8"
)

// within fails the test if f does not return within 2 seconds.
func within(t *testing.T, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2s", what)
	}
}

func TestCountLetters(t *testing.T) {
	words := []string{"go", "gopher", "channel", "héllo", "worker", "pool", "x", "concurrency"}
	var active, peak atomic.Int32
	measure := func(w string) int {
		cur := active.Add(1)
		if cur > peak.Load() {
			peak.Store(cur)
		}
		time.Sleep(10 * time.Millisecond)
		active.Add(-1)
		return utf8.RuneCountInString(w)
	}
	var got map[string]int
	within(t, "CountLetters", func() { got = CountLetters(words, 3, measure) })
	want := map[string]int{}
	for _, w := range words {
		want[w] = utf8.RuneCountInString(w)
	}
	if !maps.Equal(got, want) {
		t.Fatalf("CountLetters = %v\nwant %v", got, want)
	}
	if p := peak.Load(); p < 2 || p > 3 {
		t.Fatalf("%d measurements ran at once, want up to 3 (one per worker)", p)
	}
}
