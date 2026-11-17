//go:build !windows

package runner

import (
	"os/exec"
	"syscall"
)

// procGroup kills a command together with everything it spawned (for
// example `go test` and the test binary it runs) by putting the command in
// its own process group.
type procGroup struct{}

func newProcGroup() *procGroup { return &procGroup{} }

func (g *procGroup) prepare(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func (g *procGroup) started(*exec.Cmd) {}

func (g *procGroup) kill(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return nil
	}
	// A negative pid signals the whole process group.
	if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
		return cmd.Process.Kill()
	}
	return nil
}

func (g *procGroup) close() {}
