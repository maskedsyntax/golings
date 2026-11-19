// execing_processes_02: Replacing the process (reference solution)
package execer

import (
	"os"
	"os/exec"
	"syscall"
)

// Exec replaces the current process with tool, passing args. It returns only on failure.
func Exec(tool string, args ...string) error {
	path, err := exec.LookPath(tool)
	if err != nil {
		return err
	}
	argv := append([]string{tool}, args...) // argv[0] is the program name
	return syscall.Exec(path, argv, os.Environ())
}
