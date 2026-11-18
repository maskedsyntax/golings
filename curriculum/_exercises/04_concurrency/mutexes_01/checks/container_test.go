package container

import (
	"sync"
	"testing"
)

func TestConcurrentInc(t *testing.T) {
	c := New()
	var wg sync.WaitGroup
	for _, name := range []string{"a", "a", "b"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 10000 {
				c.Inc(name)
			}
		}()
	}
	wg.Add(1)
	go func() { // a concurrent reader
		defer wg.Done()
		for range 1000 {
			_ = c.Get("a")
		}
	}()
	wg.Wait()
	if a, b := c.Get("a"), c.Get("b"); a != 20000 || b != 10000 {
		t.Fatalf("a = %d, b = %d; want 20000 and 10000", a, b)
	}
}
