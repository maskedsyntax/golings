// spawning_processes_01: Running a command (reference solution)
package run

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
)

// Run executes name with args and returns its stdout and exit status.
func Run(ctx context.Context, name string, args ...string) (stdout string, code int, err error) {
	cmd := exec.CommandContext(ctx, name, args...) // killed when ctx ends
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if ctx.Err() != nil {
		return out.String(), -1, fmt.Errorf("%s: %w", name, ctx.Err())
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return out.String(), exitErr.ExitCode(), nil // it ran; it just failed
	}
	if err != nil {
		return "", -1, err
	}
	return out.String(), 0, nil
}
