package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("GOLINGS_RUN_MAIN") == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestProgram(t *testing.T) {
	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "GOLINGS_RUN_MAIN=1")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	code := 0
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		code = ee.ExitCode()
	} else if err != nil {
		t.Fatalf("could not run the program: %v", err)
	}
	want := "checking 3 files...\nfound 2 problems\n"
	if out.String() != want || code != 2 {
		t.Fatalf("the program printed %q and exited %d; want %q and 2", out.String(), code, want)
	}
}
