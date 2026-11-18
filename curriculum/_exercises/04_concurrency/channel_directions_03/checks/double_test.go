package double

import (
	"slices"
	"testing"
	"time"
)

func TestDoubleAll(t *testing.T) {
	done := make(chan []int, 1)
	go func() { done <- DoubleAll([]int{1, 2, 3}) }()
	select {
	case got := <-done:
		if want := []int{2, 4, 6}; !slices.Equal(got, want) {
			t.Fatalf("DoubleAll = %v, want %v", got, want)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("DoubleAll hung: no result ever arrived on results")
	}
}
