package graceful

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if os.Getenv("GOLINGS_SIGNAL_CHILD") == "1" {
		sig := Run(func() { fmt.Println("ready") }, func() { fmt.Println("cleanup") })
		fmt.Printf("exited cleanly after %v\n", sig)
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func signalChild(t *testing.T, sig os.Signal) {
	t.Helper()
	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "GOLINGS_SIGNAL_CHILD=1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { cmd.Process.Kill() })

	lines := make(chan string, 10)
	go func() {
		sc := bufio.NewScanner(stdout)
		for sc.Scan() {
			lines <- sc.Text()
		}
		close(lines)
	}()
	select {
	case l := <-lines:
		if l != "ready" {
			t.Fatalf("child printed %q before ready", l)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("the child never called ready")
	}
	if err := cmd.Process.Signal(sig); err != nil {
		t.Fatal(err)
	}

	var got []string
	timeout := time.After(5 * time.Second)
collect:
	for {
		select {
		case l, ok := <-lines:
			if !ok {
				break collect
			}
			got = append(got, l)
		case <-timeout:
			t.Fatalf("the child did not exit after %v", sig)
		}
	}
	err = cmd.Wait()
	output := strings.Join(got, "\n")
	if err != nil {
		t.Fatalf("after %v the child died (%v) instead of cleaning up; output: %q", sig, err, output)
	}
	if !strings.Contains(output, "cleanup") || !strings.Contains(output, "exited cleanly") {
		t.Fatalf("after %v the child printed %q; want cleanup to run and Run to return", sig, output)
	}
}

func TestSIGTERM(t *testing.T) { signalChild(t, syscall.SIGTERM) }

func TestSIGINT(t *testing.T) { signalChild(t, os.Interrupt) }
