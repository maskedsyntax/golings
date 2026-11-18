package max

import (
	"math/rand/v2"
	"sync"
	"testing"
)

func TestObserveSequential(t *testing.T) {
	var m Max
	if m.Value() != 0 {
		t.Fatalf("zero Max: Value = %d, want 0", m.Value())
	}
	for _, x := range []int64{5, 3, 9, 1} {
		m.Observe(x)
	}
	if got := m.Value(); got != 9 {
		t.Fatalf("after observing 5, 3, 9, 1: Value = %d, want 9", got)
	}
}

func TestObserveConcurrent(t *testing.T) {
	for range 20 {
		var m Max
		nums := rand.Perm(5000)
		var wg sync.WaitGroup
		for w := range 8 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := w; i < len(nums); i += 8 {
					m.Observe(int64(nums[i]))
				}
			}()
		}
		wg.Wait()
		if got := m.Value(); got != 4999 {
			t.Fatalf("Value = %d, want 4999: a larger value was overwritten by a smaller one", got)
		}
	}
}
