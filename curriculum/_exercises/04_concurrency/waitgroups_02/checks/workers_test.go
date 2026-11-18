package workers

import (
	"slices"
	"testing"
	"time"
)

func TestRunWorkers(t *testing.T) {
	done := make(chan []int, 1)
	go func() { done <- RunWorkers(5) }()
	select {
	case got := <-done:
		if want := []int{0, 1, 4, 9, 16}; !slices.Equal(got, want) {
			t.Fatalf("RunWorkers(5) = %v, want %v", got, want)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("RunWorkers hung: the workers' Done calls never reached the WaitGroup it waits on")
	}
}
