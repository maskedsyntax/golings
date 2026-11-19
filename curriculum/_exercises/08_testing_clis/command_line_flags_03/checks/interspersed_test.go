package interspersed

import (
	"flag"
	"io"
	"slices"
	"testing"
)

func newFS() (*flag.FlagSet, *bool, *string) {
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs, fs.Bool("v", false, ""), fs.String("o", "", "")
}

func TestParseInterspersed(t *testing.T) {
	for _, tt := range []struct {
		args     []string
		operands []string
		v        bool
		o        string
	}{
		{[]string{"a.txt", "-v", "b.txt", "-o", "out"}, []string{"a.txt", "b.txt"}, true, "out"},
		{[]string{"-v", "a"}, []string{"a"}, true, ""},
		{[]string{"--", "a", "-v"}, []string{"a", "-v"}, false, ""},
		{[]string{"x", "--", "-o", "y"}, []string{"x", "-o", "y"}, false, ""},
		{[]string{"-", "-v"}, []string{"-"}, true, ""},
		{nil, nil, false, ""},
	} {
		fs, v, o := newFS()
		operands, err := ParseInterspersed(fs, tt.args)
		if err != nil {
			t.Errorf("ParseInterspersed(%q) returned an error: %v", tt.args, err)
			continue
		}
		if !slices.Equal(operands, tt.operands) || *v != tt.v || *o != tt.o {
			t.Errorf("ParseInterspersed(%q): operands %q, -v %v, -o %q; want %q, %v, %q",
				tt.args, operands, *v, *o, tt.operands, tt.v, tt.o)
		}
	}
}

func TestParseInterspersedError(t *testing.T) {
	fs, _, _ := newFS()
	if _, err := ParseInterspersed(fs, []string{"a", "-bogus"}); err == nil {
		t.Error("an unknown flag after an operand should be an error")
	}
}
