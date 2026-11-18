package conn

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
)

func TestCloseTwice(t *testing.T) {
	var releases atomic.Int32
	c := New(func() { releases.Add(1) })
	if err := c.Close(); err != nil {
		t.Fatalf("first Close = %v, want nil", err)
	}
	if err := c.Close(); !errors.Is(err, ErrClosed) {
		t.Fatalf("second Close = %v, want ErrClosed", err)
	}
	if releases.Load() != 1 {
		t.Fatalf("release ran %d times, want 1", releases.Load())
	}
}

func TestConcurrentClose(t *testing.T) {
	for range 50 {
		var releases, nils atomic.Int32
		c := New(func() { releases.Add(1) })
		start := make(chan struct{})
		var wg sync.WaitGroup
		for range 20 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-start
				if c.Close() == nil {
					nils.Add(1)
				}
			}()
		}
		close(start)
		wg.Wait()
		if releases.Load() != 1 || nils.Load() != 1 {
			t.Fatalf("20 concurrent Close calls: release ran %d times and %d calls returned nil; want 1 and 1", releases.Load(), nils.Load())
		}
	}
}
