package describe

import (
	"testing"
	"time"
)

type color int

func (c color) String() string { return [...]string{"red", "green"}[c] }

func TestDescribe(t *testing.T) {
	for _, tt := range []struct {
		x    any
		want string
	}{
		{nil, "nil"},
		{42, "int 42"},
		{-1, "int -1"},
		{"hi", `string "hi"`},
		{`a"b`, `string "a\"b"`},
		{[]int{1, 2, 3}, "3 ints"},
		{[]int(nil), "0 ints"},
		{color(1), "stringer green"},
		{2 * time.Second, "stringer 2s"},
		{3.5, "unknown"},
		{[]string{"x"}, "unknown"},
		{int64(7), "unknown"},
	} {
		if got := Describe(tt.x); got != tt.want {
			t.Errorf("Describe(%#v) = %q, want %q", tt.x, got, tt.want)
		}
	}
}
