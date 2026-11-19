package exitcode

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestExitCode(t *testing.T) {
	for _, tt := range []struct {
		err  error
		want int
	}{
		{nil, 0},
		{errors.New("disk full"), 1},
		{&UsageError{"missing NAME"}, 2},
		{fmt.Errorf("parsing flags: %w", &UsageError{"bad -n"}), 2},
		{ErrNotFound, 3},
		{fmt.Errorf("loading config.json: %w", ErrNotFound), 3},
		{fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", ErrNotFound)), 3},
		{context.Canceled, 130},
		{fmt.Errorf("download: %w", context.Canceled), 130},
		{errors.Join(errors.New("x"), ErrNotFound), 3},
	} {
		if got := ExitCode(tt.err); got != tt.want {
			t.Errorf("ExitCode(%v) = %d, want %d", tt.err, got, tt.want)
		}
	}
}
