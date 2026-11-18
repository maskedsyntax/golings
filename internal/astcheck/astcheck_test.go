package astcheck

import (
	"strings"
	"testing"
)

const src = `package x

import str "strings"

const Limit = 3

type T struct{}

func (t *T) Close() { defer func() {}() }

func Loop(xs []int) (n int) {
	for _, x := range xs {
		switch {
		case x > 0:
			n++
		}
	}
	return str.ToUpper("a") == "A" && n > 0 || true
}

func Pipe(in <-chan int, out chan<- int) {}
func Both(c chan int)                    {}
`

func TestRules(t *testing.T) {
	files := map[string][]byte{"x.go": []byte(src)}
	for _, tt := range []struct {
		rule string
		ok   bool
	}{
		{"uses_defer:T.Close", true},
		{"uses_defer:Loop", false},
		{"uses_range:Loop", true},
		{"uses_switch:Loop", true},
		{"uses_switch:Pipe", false},
		{"declares_const:Limit", true},
		{"declares_const:Max", false},
		{"calls:strings.ToUpper", true},
		{"calls:strings.ToLower", false},
		{"no_call:strings.ToUpper", false},
		{"no_call:sort.Ints", true},
		{"chan_params_directional:Pipe", true},
		{"chan_params_directional:Both", false},
		{"uses_defer:Missing", false},
	} {
		v, err := Check(files, []string{tt.rule})
		if err != nil {
			t.Fatalf("%s: %v", tt.rule, err)
		}
		if ok := len(v) == 0; ok != tt.ok {
			t.Errorf("%s: satisfied=%v, want %v (%v)", tt.rule, ok, tt.ok, v)
		}
	}
}

func TestParse(t *testing.T) {
	if _, err := Parse("uses_defer:F"); err != nil {
		t.Error(err)
	}
	for _, bad := range []string{"rm -rf", "uses_defer", "calls:Println"} {
		if _, err := Parse(bad); err == nil {
			t.Errorf("Parse(%q) succeeded", bad)
		}
	}
}

func TestSyntaxErrorIsError(t *testing.T) {
	_, err := Check(map[string][]byte{"x.go": []byte("package x\nfunc {")}, []string{"uses_defer:F"})
	if err == nil || !strings.Contains(err.Error(), "x.go") {
		t.Errorf("err = %v", err)
	}
}
