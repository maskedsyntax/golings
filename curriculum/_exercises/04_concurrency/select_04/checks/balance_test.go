package balance

import (
	"slices"
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

func feed(n int) <-chan int {
	in := make(chan int)
	go func() {
		for i := range n {
			in <- i
		}
	}()
	return in
}

func TestBalanceAvoidsStuckConsumer(t *testing.T) {
	a := make(chan int) // nobody ever receives on a
	b := make(chan int)
	var got []int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 5 {
			select {
			case v := <-b:
				got = append(got, v)
			case <-time.After(2 * time.Second):
				return
			}
		}
	}()
	within(t, "Balance with a stuck consumer on a", func() { Balance(feed(5), a, b, 5) })
	wg.Wait()
	if want := []int{0, 1, 2, 3, 4}; !slices.Equal(got, want) {
		t.Fatalf("b received %v, want %v", got, want)
	}
}

func TestBalanceUsesBoth(t *testing.T) {
	a, b := make(chan int), make(chan int)
	var mu sync.Mutex
	var got []int
	var wg sync.WaitGroup
	for _, ch := range []chan int{a, b} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case v := <-ch:
					mu.Lock()
					got = append(got, v)
					mu.Unlock()
				case <-time.After(300 * time.Millisecond):
					return
				}
			}
		}()
	}
	within(t, "Balance", func() { Balance(feed(20), a, b, 20) })
	wg.Wait()
	slices.Sort(got)
	if len(got) != 20 || got[0] != 0 || got[19] != 19 {
		t.Fatalf("consumers received %v, want 0..19", got)
	}
}
