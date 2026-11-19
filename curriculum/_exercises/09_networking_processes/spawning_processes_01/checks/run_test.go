package run

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// The test binary doubles as the child program: with GOLINGS_HELPER set, it
// behaves like a small command-line tool instead of running tests.
func TestMain(m *testing.M) {
	switch os.Getenv("GOLINGS_HELPER") {
	case "":
		os.Exit(m.Run())
	case "echo":
		fmt.Println(strings.Join(os.Args[1:], " "))
		os.Exit(0)
	case "fail":
		fmt.Println("partial")
		fmt.Fprintln(os.Stderr, "boom: something broke")
		os.Exit(3)
	case "sleep":
		time.Sleep(10 * time.Second)
		os.Exit(0)
	case "upper":
		data, _ := io.ReadAll(os.Stdin)
		fmt.Print(strings.ToUpper(string(data)))
		os.Exit(0)
	}
	os.Exit(99)
}

// helper makes os.Args[0] behave as the named helper program in child processes.
func helper(t *testing.T, mode string) string {
	t.Setenv("GOLINGS_HELPER", mode)
	return os.Args[0]
}

func TestRunOK(t *testing.T) {
	out, code, err := Run(context.Background(), helper(t, "echo"), "hello", "world")
	if err != nil || code != 0 || out != "hello world\n" {
		t.Fatalf("Run(echo hello world) = %q, %d, %v; want \"hello world\\n\", 0, nil", out, code, err)
	}
}

func TestRunExitStatus(t *testing.T) {
	out, code, err := Run(context.Background(), helper(t, "fail"))
	if err != nil || code != 3 || out != "partial\n" {
		t.Fatalf("Run(fail) = %q, %d, %v; want \"partial\\n\", 3, nil (a non-zero status is not an error here)", out, code, err)
	}
}

func TestRunMissing(t *testing.T) {
	if _, _, err := Run(context.Background(), "golings-no-such-program-xyz"); err == nil {
		t.Fatal("Run of a program that doesn't exist returned no error")
	}
}

func TestRunTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, _, err := Run(ctx, helper(t, "sleep"))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Run error = %v, want one wrapping context.DeadlineExceeded", err)
	}
	if d := time.Since(start); d > 5*time.Second {
		t.Fatalf("Run took %v; the process should be killed when ctx expires", d)
	}
}
