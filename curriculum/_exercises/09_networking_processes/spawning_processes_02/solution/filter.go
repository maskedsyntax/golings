// spawning_processes_02: Pipes and stderr (reference solution)
package filter

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Filter runs name with args, writes input to its stdin and returns its stdout.
func Filter(ctx context.Context, input, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdin = strings.NewReader(input) // exec closes stdin after copying it
	var out, errOut bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errOut
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s: %w: %s", name, err, strings.TrimSpace(errOut.String()))
	}
	return out.String(), nil
}
