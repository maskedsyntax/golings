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

func TestCallsAreIndependent(t *testing.T) {
	steps := []struct {
		args []string
		want string
	}{
		{[]string{"greet", "-loud", "ada"}, "HELLO, ADA\n"},
		{[]string{"greet", "bob"}, "hello, bob\n"},
		{[]string{"count", "-by", "5"}, "0\n5\n10\n"},
		{[]string{"count"}, "0\n1\n2\n"},
		{[]string{"greet", "-loud", "cy"}, "HELLO, CY\n"},
		{[]string{"greet", "di"}, "hello, di\n"},
	}
	for _, s := range steps {
		out, errOut, code := call(s.args...)
		if out != s.want || code != 0 {
			t.Errorf("%s: stdout %q, stderr %q, exit %d; want %q and 0", strings.Join(s.args, " "), out, errOut, code, s.want)
		}
	}
}

func TestNoArgs(t *testing.T) {
	out, errOut, code := call()
	if code != 2 || out != "" || errOut == "" {
		t.Errorf("no arguments: stdout %q, stderr %q, exit %d; want a usage message on stderr and exit 2", out, errOut, code)
	}
}

func TestUnknown(t *testing.T) {
	if _, errOut, code := call("dance"); code != 2 || !strings.Contains(errOut, "dance") {
		t.Errorf("unknown subcommand: stderr %q, exit %d; want a message naming it and exit 2", errOut, code)
	}
}
