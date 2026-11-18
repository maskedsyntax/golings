package loops

import (
	"reflect"
	"testing"
)

func TestCountDown(t *testing.T) {
	if got := CountDown(3); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("CountDown(3) = %v, want [3 2 1]", got)
	}
	for _, n := range []int{0, -5} {
		if got := CountDown(n); len(got) != 0 {
			t.Errorf("CountDown(%d) = %v, want no elements", n, got)
		}
	}
}

func TestSumTo(t *testing.T) {
	for _, tt := range []struct{ n, want int }{{4, 10}, {1, 1}, {0, 0}, {-3, 0}, {100, 5050}} {
		if got := SumTo(tt.n); got != tt.want {
			t.Errorf("SumTo(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
