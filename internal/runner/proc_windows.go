//go:build windows

package runner

import (
	"os/exec"
	"unsafe"

	"golang.org/x/sys/windows"
)

// procGroup kills a command and all of its descendants by assigning it to a
// Job Object configured to kill every member when the job is closed.
type procGroup struct {
	job windows.Handle
}

func newProcGroup() *procGroup {
	g := &procGroup{}
	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		return g
	}
	info := windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION{
		BasicLimitInformation: windows.JOBOBJECT_BASIC_LIMIT_INFORMATION{
			LimitFlags: windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE,
		},
	}
	if _, err := windows.SetInformationJobObject(job, windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info))); err != nil {
		windows.CloseHandle(job)
		return g
	}
	g.job = job
	return g
}

func (g *procGroup) prepare(*exec.Cmd) {}

// started assigns the process to the job. Processes spawned before the
// assignment completes are not tracked; for `go test` the test binary starts
// well after this point.
func (g *procGroup) started(cmd *exec.Cmd) {
	if g.job == 0 || cmd.Process == nil {
		return
	}
	h, err := windows.OpenProcess(windows.PROCESS_SET_QUOTA|windows.PROCESS_TERMINATE, false, uint32(cmd.Process.Pid))
	if err != nil {
		return
	}
	defer windows.CloseHandle(h)
	_ = windows.AssignProcessToJobObject(g.job, h)
}

func (g *procGroup) kill(cmd *exec.Cmd) error {
	if g.job != 0 {
		if err := windows.TerminateJobObject(g.job, 1); err == nil {
			return nil
		}
	}
	if cmd.Process != nil {
		return cmd.Process.Kill()
	}
	return nil
}

// close releases the job; KILL_ON_JOB_CLOSE terminates any leftovers.
func (g *procGroup) close() {
	if g.job != 0 {
		windows.CloseHandle(g.job)
		g.job = 0
	}
}
