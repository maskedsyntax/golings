package ranges

import (
	"fmt"
	"slices"
	"testing"
)

func TestSquares(t *testing.T) {
	if got := Squares(4); !slices.Equal(got, []int{1, 4, 9, 16}) {
		t.Errorf("Squares(4) = %v, want [1 4 9 16]", got)
	}
	if got := Squares(1); !slices.Equal(got, []int{1}) {
		t.Errorf("Squares(1) = %v, want [1]", got)
	}
	for _, n := range []int{0, -3} {
		if got := Squares(n); len(got) != 0 {
			t.Errorf("Squares(%d) = %v, want empty", n, got)
		}
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[string]int{}
	var want []string
	for i := range 26 {
		k := fmt.Sprintf("%c%d", 'a'+i, i)
		m[k] = i
		want = append(want, k)
	}
	for range 5 { // map order changes between iterations; one lucky run must not pass
		if got := SortedKeys(m); !slices.Equal(got, want) {
			t.Fatalf("SortedKeys returned %v, want %v", got, want)
		}
	}
	if got := SortedKeys(nil); len(got) != 0 {
		t.Errorf("SortedKeys(nil) = %v, want empty", got)
	}
}
