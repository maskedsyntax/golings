package supervise

import (
	"os"
	"syscall"
	"testing"
	"time"
)

// within fails the test if f doesn't return in time, instead of hanging.
func within(t *testing.T, what string, d time.Duration, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	select {
	case <-done:
	case <-time.After(d):
		t.Fatalf("%s did not finish within %v", what, d)
	}
}

func feed(sigs ...os.Signal) <-chan os.Signal {
	ch := make(chan os.Signal, len(sigs))
	for _, s := range sigs {
		ch <- s
	}
	return ch
}

func TestSupervise(t *testing.T) {
	for _, tt := range []struct {
		name     string
		sigs     []os.Signal
		reloads  int
		stopWith os.Signal
	}{
		{"term after reloads", []os.Signal{syscall.SIGHUP, syscall.SIGHUP, syscall.SIGTERM}, 2, syscall.SIGTERM},
		{"interrupt", []os.Signal{os.Interrupt}, 0, os.Interrupt},
		{"ignores others", []os.Signal{syscall.SIGALRM, syscall.SIGHUP, os.Interrupt, syscall.SIGHUP}, 1, os.Interrupt},
	} {
		reloads, shutdowns := 0, 0
		var got os.Signal
		within(t, tt.name, time.Second, func() {
			got = Supervise(feed(tt.sigs...), func() { reloads++ }, func() { shutdowns++ })
		})
		if got != tt.stopWith || reloads != tt.reloads || shutdowns != 1 {
			t.Errorf("%s: returned %v after %d reloads and %d shutdowns; want %v, %d, 1",
				tt.name, got, reloads, shutdowns, tt.stopWith, tt.reloads)
		}
	}
}

func TestSuperviseClosed(t *testing.T) {
	ch := make(chan os.Signal, 2)
	ch <- syscall.SIGHUP
	close(ch)
	reloads, shutdowns := 0, 0
	var got os.Signal
	within(t, "Supervise on a closed channel", time.Second, func() {
		got = Supervise(ch, func() { reloads++ }, func() { shutdowns++ })
	})
	if got != nil || reloads != 1 || shutdowns != 0 {
		t.Errorf("returned %v after %d reloads and %d shutdowns; want nil, 1, 0", got, reloads, shutdowns)
	}
}
