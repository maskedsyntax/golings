// Mutant: a non-zero exit is returned as an error.
package run

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func Run(ctx context.Context, name string, args ...string) (string, int, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		if ctx.Err() != nil {
			return "", -1, fmt.Errorf("%w", ctx.Err())
		}
		return out.String(), -1, err
	}
	return out.String(), 0, nil
}
