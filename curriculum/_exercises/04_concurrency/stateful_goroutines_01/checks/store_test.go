package store

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

func TestSetThenGet(t *testing.T) {
	s := NewStore()
	within(t, "Set/Get", func() {
		s.Set("a", 1)
		s.Set("b", 2)
		s.Set("a", 3)
	})
	var a, b, c int
	within(t, "Get", func() { a, b, c = s.Get("a"), s.Get("b"), s.Get("missing") })
	if a != 3 || b != 2 || c != 0 {
		t.Fatalf("Get a, b, missing = %d, %d, %d; want 3, 2, 0", a, b, c)
	}
}

func TestConcurrentUse(t *testing.T) {
	s := NewStore()
	within(t, "concurrent Set/Get", func() {
		var wg sync.WaitGroup
		for w := range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				key := fmt.Sprint("k", w)
				for i := 1; i <= 100; i++ {
					s.Set(key, i)
					if got := s.Get(key); got != i {
						t.Errorf("worker %d: Get after Set(%d) = %d", w, i, got)
						return
					}
				}
			}()
		}
		wg.Wait()
	})
}
