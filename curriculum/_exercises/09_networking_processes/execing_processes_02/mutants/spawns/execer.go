// Mutant: runs a child process and exits instead of exec'ing.
package execer

import (
	"os"
	"os/exec"
)

func Exec(tool string, args ...string) error {
	path, err := exec.LookPath(tool)
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
