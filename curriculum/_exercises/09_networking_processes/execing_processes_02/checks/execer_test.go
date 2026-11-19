package execer

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

// The test binary execs itself: stage "before" prints its PID and calls Exec;
// stage "after" is the replacement program and prints what it received.
func TestMain(m *testing.M) {
	switch os.Getenv("GOLINGS_EXEC_STAGE") {
	case "before":
		fmt.Printf("before pid=%d\n", os.Getpid())
		os.Setenv("GOLINGS_EXEC_STAGE", "after")
		err := Exec(os.Args[0], "alpha", "beta")
		fmt.Printf("Exec returned: %v\n", err)
		os.Exit(1)
	case "after":
		fmt.Printf("after pid=%d args=%q\n", os.Getpid(), os.Args[1:])
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestExec(t *testing.T) {
	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "GOLINGS_EXEC_STAGE=before")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("the child failed: %v\n%s", err, out)
	}
	before := regexp.MustCompile(`before pid=(\d+)`).FindSubmatch(out)
	after := regexp.MustCompile(`after pid=(\d+) args=(.*)`).FindSubmatch(out)
	if before == nil || after == nil {
		t.Fatalf("unexpected output:\n%s", out)
	}
	if string(before[1]) != string(after[1]) {
		t.Errorf("the new program ran as PID %s, not %s: it was spawned, not exec'd", after[1], before[1])
	}
	if got, want := string(after[2]), `["alpha" "beta"]`; got != want {
		t.Errorf("the new program received arguments %s, want %s", got, want)
	}
}

func TestExecMissing(t *testing.T) {
	if err := Exec("golings-no-such-program-xyz"); err == nil {
		t.Fatal("Exec of a missing program returned no error")
	}
}
