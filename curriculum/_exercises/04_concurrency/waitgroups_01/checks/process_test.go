package process

import (
	"testing"
	"time"
)

func TestProcessAllWaits(t *testing.T) {
	for range 3 {
		done := make([]bool, 10)
		ProcessAll(10, func(i int) {
			time.Sleep(10 * time.Millisecond)
			done[i] = true
		})
		for i, d := range done {
			if !d {
				t.Fatalf("ProcessAll returned before work(%d) finished", i)
			}
		}
	}
}
