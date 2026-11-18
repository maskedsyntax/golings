package ranges

import (
	"maps"
	"testing"
)

func TestSum(t *testing.T) {
	for _, tt := range []struct {
		nums []int
		want int
	}{{[]int{5, 6, 7}, 18}, {[]int{-1, 1}, 0}, {nil, 0}, {[]int{10}, 10}} {
		if got := Sum(tt.nums); got != tt.want {
			t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestCountRunesASCII(t *testing.T) {
	if got, want := CountRunes("banana"), map[rune]int{'b': 1, 'a': 3, 'n': 2}; !maps.Equal(got, want) {
		t.Errorf("CountRunes(%q) = %v, want %v", "banana", got, want)
	}
}

func TestCountRunesUnicode(t *testing.T) {
	for _, tt := range []struct {
		s    string
		want map[rune]int
	}{
		{"héé", map[rune]int{'h': 1, 'é': 2}},
		{"日本日", map[rune]int{'日': 2, '本': 1}},
		{"", map[rune]int{}},
	} {
		if got := CountRunes(tt.s); !maps.Equal(got, tt.want) {
			t.Errorf("CountRunes(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
