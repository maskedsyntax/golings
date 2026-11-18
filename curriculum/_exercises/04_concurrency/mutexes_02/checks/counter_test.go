package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Fatalf("Value = %d after 1000 Inc calls, want 1000", got)
	}
}
