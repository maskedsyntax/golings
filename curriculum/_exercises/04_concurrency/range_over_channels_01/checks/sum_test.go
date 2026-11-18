package sum

import (
	"testing"
	"time"
)

func send(nums ...int) <-chan int {
	ch := make(chan int, len(nums))
	for _, n := range nums {
		ch <- n
	}
	close(ch)
	return ch
}

func TestSum(t *testing.T) {
	for _, tt := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{4, 5, 6, 7, 8}, 30},
		{[]int{9}, 9},
		{nil, 0},
	} {
		done := make(chan int, 1)
		go func() { done <- Sum(send(tt.nums...)) }()
		select {
		case got := <-done:
			if got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		case <-time.After(2 * time.Second):
			t.Fatalf("Sum(%v) blocked", tt.nums)
		}
	}
}
