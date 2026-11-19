package filter

import (
	"context"
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

func TestFilter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	got, err := Filter(ctx, "hello\npipes\n", helper(t, "upper"))
	if err != nil || got != "HELLO\nPIPES\n" {
		t.Fatalf("Filter = %q, %v; want \"HELLO\\nPIPES\\n\", nil", got, err)
	}
}

func TestFilterLargeInput(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	in := strings.Repeat("abc\n", 256*1024) // 1 MiB, more than any pipe buffer
	got, err := Filter(ctx, in, helper(t, "upper"))
	if err != nil || got != strings.ToUpper(in) {
		t.Fatalf("Filter(1 MiB) returned %d bytes, %v", len(got), err)
	}
}

func TestFilterStderr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := Filter(ctx, "", helper(t, "fail"))
	if err == nil {
		t.Fatal("Filter returned no error for a command that exited with status 3")
	}
	if !strings.Contains(err.Error(), "boom: something broke") {
		t.Fatalf("error %q should include the command's stderr", err)
	}
}
