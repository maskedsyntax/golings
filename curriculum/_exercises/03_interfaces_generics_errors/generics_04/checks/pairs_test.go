package pairs

import (
	"reflect"
	"slices"
	"testing"
)

func TestEntries(t *testing.T) {
	m := map[string]int{}
	var want []Pair[string, int]
	for i, k := range []string{"apple", "banana", "cherry", "date", "elder", "fig", "grape", "honeydew"} {
		m[k] = i
		want = append(want, Pair[string, int]{k, i})
	}
	for range 5 {
		if got := Entries(m); !slices.Equal(got, want) {
			t.Fatalf("Entries = %v, want %v", got, want)
		}
	}
	nums := Entries(map[float64]bool{2.5: true, -1: false, 0: true})
	if want := []Pair[float64, bool]{{-1, false}, {0, true}, {2.5, true}}; !slices.Equal(nums, want) {
		t.Errorf("Entries(float64 keys) = %v, want %v", nums, want)
	}
	if got := Entries(map[int]int{}); len(got) != 0 {
		t.Errorf("Entries(empty) = %v", got)
	}
}

func TestGroupBy(t *testing.T) {
	words := []string{"go", "c", "js", "rust", "d", "zig"}
	got := GroupBy(words, func(s string) int { return len(s) })
	want := map[int][]string{1: {"c", "d"}, 2: {"go", "js"}, 3: {"zig"}, 4: {"rust"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy(by length) = %v, want %v", got, want)
	}
	parity := GroupBy([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	if !reflect.DeepEqual(parity, map[bool][]int{false: {1, 3, 5}, true: {2, 4}}) {
		t.Errorf("GroupBy(even) = %v", parity)
	}
}
