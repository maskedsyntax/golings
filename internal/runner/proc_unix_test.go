//go:build !windows

package runner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"
)

// The test binary doubles as a helper process: with GOLINGS_TEST_HELPER set
// it either spawns a sleeping grandchild (and reports its PID) or sleeps.
func TestMain(m *testing.M) {
	switch os.Getenv("GOLINGS_TEST_HELPER") {
	case "spawn":
		child := exec.Command(os.Args[0])
		child.Env = append(os.Environ(), "GOLINGS_TEST_HELPER=sleep")
		if err := child.Start(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", child.Process.Pid)
		os.Stdout.Sync()
		time.Sleep(time.Hour)
		os.Exit(0)
	case "sleep":
		time.Sleep(time.Hour)
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func alive(pid int) bool {
	return syscall.Kill(pid, 0) == nil
}

// TestCancelKillsProcessTree checks that a timeout kills the grandchild too,
// not just the direct child (the `go test` -> test binary situation).
func TestCancelKillsProcessTree(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	env := append(os.Environ(), "GOLINGS_TEST_HELPER=spawn")
	start := time.Now()
	res := execCommand(ctx, t.TempDir(), env, outputLimit, os.Args[0])
	if !res.TimedOut {
		t.Fatalf("result = %+v, want timed out", res)
	}
	if d := time.Since(start); d > 10*time.Second {
		t.Errorf("cancellation took %v", d)
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(bytes.Split(res.Stdout, []byte("\n"))[0])))
	if err != nil {
		t.Fatalf("could not read grandchild pid from %q", res.Stdout)
	}
	deadline := time.Now().Add(5 * time.Second)
	for alive(pid) && time.Now().Before(deadline) {
		// The grandchild becomes a zombie until reaped by init; signal 0 can
		// still succeed for zombies, so also check it is no longer running.
		if out, err := exec.Command("ps", "-o", "stat=", "-p", strconv.Itoa(pid)).Output(); err != nil || strings.HasPrefix(strings.TrimSpace(string(out)), "Z") {
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
	if alive(pid) {
		syscall.Kill(pid, syscall.SIGKILL)
		t.Errorf("grandchild %d survived cancellation", pid)
	}
}
