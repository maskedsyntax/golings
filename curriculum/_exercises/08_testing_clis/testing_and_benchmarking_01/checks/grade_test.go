package intmin

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

// golings grader: your tests must pass against the real IntMin and fail
// against each broken version below.

type bug struct {
	name, desc string
	fn         func(a, b int) int
}

var bugs = []bug{
	{"second", "IntMin always returns its second argument", func(a, b int) int { return b }},
	{"max", "IntMin returns the larger number", func(a, b int) int { return max(a, b) }},
}

func TestMain(m *testing.M) {
	if name := os.Getenv("GOLINGS_BUG"); name != "" {
		for _, b := range bugs {
			if b.name == name {
				IntMin = b.fn
			}
		}
	}
	os.Exit(m.Run())
}

func runYourTests(t *testing.T, env string, args ...string) ([]byte, bool) {
	t.Helper()
	args = append([]string{"-test.run=^Test", "-test.skip=^TestYour", "-test.count=1"}, args...)
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), env)
	out, err := cmd.CombinedOutput()
	var ee *exec.ExitError
	switch {
	case err == nil:
		return out, true
	case errors.As(err, &ee):
		return out, false
	}
	t.Fatalf("could not run your tests: %v\n%s", err, out)
	return nil, false
}

func TestYourTestsCatchBugs(t *testing.T) {
	if os.Getenv("GOLINGS_BUG") != "" {
		t.Skip("running inside a bug check")
	}
	for _, b := range bugs {
		if _, passed := runYourTests(t, "GOLINGS_BUG="+b.name); passed {
			t.Errorf("your tests still pass when %s; add a case that catches it", b.desc)
		}
	}
}
