package describe

import "testing"

func TestDescribe(t *testing.T) {
	var nilSlice []int
	for _, tt := range []struct {
		name string
		in   any
		want string
	}{
		{"nil", nil, "nil"},
		{"int", 42, "int 42"},
		{"negative int", -1, "int -1"},
		{"string", "hi", `string "hi"`},
		{"empty string", "", `string ""`},
		{"bool", true, "bool true"},
		{"slice", []int{1, 2, 3}, "[]int of length 3"},
		{"nil slice is still a []int", nilSlice, "[]int of length 0"},
		{"int64 is not int", int64(5), "unknown"},
		{"float", 1.5, "unknown"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Describe(tt.in); got != tt.want {
				t.Errorf("Describe(%#v) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
