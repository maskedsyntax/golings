package cache

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestComputeOncePerKey(t *testing.T) {
	c := New()
	var computes atomic.Int32
	compute := func() string {
		computes.Add(1)
		time.Sleep(20 * time.Millisecond)
		return "value"
	}
	start := make(chan struct{})
	var wg sync.WaitGroup
	results := make([]string, 30)
	for i := range results {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			results[i] = c.Get("k", compute)
		}()
	}
	close(start)
	wg.Wait()
	if n := computes.Load(); n != 1 {
		t.Fatalf("compute ran %d times for one key, want 1", n)
	}
	for i, r := range results {
		if r != "value" {
			t.Fatalf("result %d = %q, want %q", i, r, "value")
		}
	}
}

func TestDifferentKeys(t *testing.T) {
	c := New()
	var computes atomic.Int32
	mk := func(v string) func() string { return func() string { computes.Add(1); return v } }
	if c.Get("a", mk("1")) != "1" || c.Get("b", mk("2")) != "2" || c.Get("a", mk("x")) != "1" {
		t.Fatal("Get returned the wrong value")
	}
	if computes.Load() != 2 {
		t.Fatalf("compute ran %d times for 2 keys, want 2", computes.Load())
	}
}
