// Mutant: input fixed, stderr still discarded.
package filter

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)

func Filter(ctx context.Context, input, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}
