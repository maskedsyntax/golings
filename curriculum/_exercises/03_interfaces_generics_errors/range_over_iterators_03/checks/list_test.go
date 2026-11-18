package list

import (
	"fmt"
	"slices"
	"testing"
)

func pairs[T any](seq func(func(int, T) bool)) []string {
	var out []string
	for i, v := range seq {
		out = append(out, fmt.Sprintf("%d:%v", i, v))
	}
	return out
}

func TestAll(t *testing.T) {
	var l List[string]
	for _, s := range []string{"a", "b", "c"} {
		l.Push(s)
	}
	if got := pairs(l.All()); !slices.Equal(got, []string{"0:a", "1:b", "2:c"}) {
		t.Errorf("All() = %v, want [0:a 1:b 2:c]", got)
	}
	var empty List[int]
	if got := pairs(empty.All()); len(got) != 0 {
		t.Errorf("All() of an empty list = %v", got)
	}
}

func TestBackward(t *testing.T) {
	var l List[int]
	for _, v := range []int{10, 20, 30} {
		l.Push(v)
	}
	if got := pairs(l.Backward()); !slices.Equal(got, []string{"2:30", "1:20", "0:10"}) {
		t.Errorf("Backward() = %v, want [2:30 1:20 0:10]", got)
	}
	var empty List[int]
	if got := pairs(empty.Backward()); len(got) != 0 {
		t.Errorf("Backward() of an empty list = %v", got)
	}
}

func TestBreak(t *testing.T) {
	var l List[int]
	for v := range 10 {
		l.Push(v)
	}
	var fwd, back []int
	for i, v := range l.All() {
		if i == 2 {
			break
		}
		fwd = append(fwd, v)
	}
	for _, v := range l.Backward() {
		if v == 7 {
			break
		}
		back = append(back, v)
	}
	if !slices.Equal(fwd, []int{0, 1}) || !slices.Equal(back, []int{9, 8}) {
		t.Errorf("with break: All saw %v (want [0 1]), Backward saw %v (want [9 8])", fwd, back)
	}
}
