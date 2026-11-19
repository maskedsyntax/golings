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

func TestSubcommands(t *testing.T) {
	for _, tt := range []struct {
		args []string
		want string
	}{
		{[]string{"add", "buy", "milk"}, "added: buy milk (priority 1)\n"},
		{[]string{"add", "-p", "3", "call", "mum"}, "added: call mum (priority 3)\n"},
		{[]string{"list"}, "listing open tasks\n"},
		{[]string{"list", "-all"}, "listing all tasks\n"},
	} {
		out, errOut, code := call(tt.args...)
		if out != tt.want || code != 0 {
			t.Errorf("todo %s: stdout %q, stderr %q, exit %d; want %q and 0", strings.Join(tt.args, " "), out, errOut, code, tt.want)
		}
	}
}

func TestUsageErrors(t *testing.T) {
	for _, tt := range []struct {
		args    []string
		mention string
	}{
		{nil, "usage"},
		{[]string{"remove", "x"}, "remove"},
		{[]string{"add"}, ""},
		{[]string{"add", "-p", "high", "x"}, ""},
		{[]string{"list", "-everything"}, ""},
	} {
		out, errOut, code := call(tt.args...)
		if code != 2 || out != "" || errOut == "" || !strings.Contains(errOut, tt.mention) {
			t.Errorf("todo %s: stdout %q, stderr %q, exit %d; want a message on stderr (mentioning %q) and exit 2",
				strings.Join(tt.args, " "), out, errOut, code, tt.mention)
		}
	}
}
