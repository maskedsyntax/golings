// Mutant: returns early but reports success.
package wait

import (
	"context"
	"time"
)

func Wait(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
	case <-ctx.Done():
	}
	return nil
}
