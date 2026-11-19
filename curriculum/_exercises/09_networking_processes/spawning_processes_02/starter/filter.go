// spawning_processes_02: Pipes and stderr
//
// Filter hangs, and its errors are unhelpful. Fix it. See README.md.
//
// Reference: https://gobyexample.com/spawning-processes
package filter

import (
	"bytes"
	"context"
	"io"
	"os/exec"
)

// Filter runs name with args, writes input to its stdin and returns its stdout.
func Filter(ctx context.Context, input, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		return "", err
	}
	io.WriteString(stdin, input)
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return out.String(), nil
}
