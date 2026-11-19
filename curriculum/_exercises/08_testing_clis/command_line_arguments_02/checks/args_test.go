package args

import (
	"slices"
	"testing"
)

func TestSplit(t *testing.T) {
	for _, tt := range []struct {
		args, flags, operands []string
	}{
		{[]string{"-v", "a.txt", "--name=x", "-", "--", "-not-a-flag", "b.txt"},
			[]string{"-v", "--name=x"}, []string{"a.txt", "-", "-not-a-flag", "b.txt"}},
		{[]string{"a", "b"}, nil, []string{"a", "b"}},
		{[]string{"-a", "-b"}, []string{"-a", "-b"}, nil},
		{[]string{"--", "--"}, nil, []string{"--"}},
		{nil, nil, nil},
	} {
		flags, operands := Split(tt.args)
		if !slices.Equal(flags, tt.flags) || !slices.Equal(operands, tt.operands) {
			t.Errorf("Split(%q) = %q, %q; want %q, %q", tt.args, flags, operands, tt.flags, tt.operands)
		}
	}
}
