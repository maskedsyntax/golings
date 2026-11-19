// execing_processes_02: Replacing the process
//
// The exec'd program receives the wrong arguments. Fix Exec. See README.md.
// (Linux and macOS only.)
//
// Reference: https://gobyexample.com/execing-processes
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
	return syscall.Exec(path, args, os.Environ())
}
