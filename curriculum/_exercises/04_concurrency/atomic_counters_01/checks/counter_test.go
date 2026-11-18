package counter

import (
	"sync"
	"testing"
)

func TestConcurrentIncrements(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 50000 {
		t.Fatalf("Value = %d after 50000 concurrent Inc calls (%d lost)", got, 50000-got)
	}
}

func TestAdd(t *testing.T) {
	var c Counter
	c.Add(10)
	c.Add(-3)
	c.Inc()
	if got := c.Value(); got != 8 {
		t.Fatalf("Value = %d, want 8", got)
	}
}
