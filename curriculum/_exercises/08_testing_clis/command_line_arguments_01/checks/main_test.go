package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// runMain runs this package's main() in a child process with the given
// command-line arguments.
func runMain(t *testing.T, args ...string) (stdout, stderr string, code int) {
	t.Helper()
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), "GOLINGS_RUN_MAIN=1")
	var out, errOut strings.Builder
	cmd.Stdout, cmd.Stderr = &out, &errOut
	err := cmd.Run()
	var ee *exec.ExitError
	switch {
	case err == nil:
	case errors.As(err, &ee):
		code = ee.ExitCode()
	default:
		t.Fatalf("could not run the program: %v", err)
	}
	return out.String(), errOut.String(), code
}

func TestMain(m *testing.M) {
	if os.Getenv("GOLINGS_RUN_MAIN") == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestProgram(t *testing.T) {
	out, errOut, code := runMain(t, "Ada", "Bob")
	if want := "Hello, Ada!\nHello, Bob!\n"; out != want || code != 0 {
		t.Errorf("greet Ada Bob printed %q and exited %d; want %q and 0", out, code, want)
	}
	if errOut != "" {
		t.Errorf("greet Ada Bob wrote to stderr: %q", errOut)
	}
}

func TestProgramNoArgs(t *testing.T) {
	out, errOut, code := runMain(t)
	if code != 2 || out != "" || !strings.Contains(errOut, "usage") {
		t.Errorf("greet with no arguments: stdout %q, stderr %q, exit %d; want only a usage message on stderr and exit 2", out, errOut, code)
	}
}
