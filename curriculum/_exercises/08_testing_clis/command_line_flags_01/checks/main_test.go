package main

import (
	"strings"
	"testing"
)

func call(args ...string) (string, string, int) {
	var out, errOut strings.Builder
	code := run(args, &out, &errOut)
	return out.String(), errOut.String(), code
}

func TestRun(t *testing.T) {
	for _, tt := range []struct {
		args []string
		want string
	}{
		{[]string{"-name", "Ada", "-n", "2"}, "Hello, Ada!\nHello, Ada!\n"},
		{[]string{"-name=Bob"}, "Hello, Bob!\n"},
		{nil, "Hello, World!\n"},
	} {
		out, errOut, code := call(tt.args...)
		if out != tt.want || code != 0 || errOut != "" {
			t.Errorf("run(%q): stdout %q, stderr %q, exit %d; want %q, \"\", 0", tt.args, out, errOut, code, tt.want)
		}
	}
}

func TestRunHelp(t *testing.T) {
	out, errOut, code := call("-h")
	if code != 0 || out != "" || !strings.Contains(errOut, "-name") {
		t.Errorf("run(-h): stdout %q, stderr %q, exit %d; want only usage on stderr and exit 0", out, errOut, code)
	}
}

func TestRunBadFlag(t *testing.T) {
	out, errOut, code := call("-bogus")
	if code != 2 || out != "" || !strings.Contains(errOut, "bogus") {
		t.Errorf("run(-bogus): stdout %q, stderr %q, exit %d; want an error on stderr and exit 2", out, errOut, code)
	}
}
