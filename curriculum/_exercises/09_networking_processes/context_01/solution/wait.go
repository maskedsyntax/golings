// context_01: Cancellable waits (reference solution)
package wait

import (
	"context"
	"time"
)

// Wait pauses for d, or until ctx is done, whichever comes first.
func Wait(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err() // Canceled or DeadlineExceeded
	}
}
