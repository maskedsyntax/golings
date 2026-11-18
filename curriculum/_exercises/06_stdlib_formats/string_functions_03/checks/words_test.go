package words

import "testing"

func TestCountWord(t *testing.T) {
	for _, tt := range []struct {
		text, word string
		want       int
	}{
		{"Go go GO! gopher, ago; go.", "go", 4},
		{"the cat sat on the mat", "the", 2},
		{"The THE tHe", "the", 3},
		{"concatenate", "cat", 0},
		{"a b", "", 0},
		{"", "go", 0},
	} {
		if got := CountWord(tt.text, tt.word); got != tt.want {
			t.Errorf("CountWord(%q, %q) = %d, want %d", tt.text, tt.word, got, tt.want)
		}
	}
}
