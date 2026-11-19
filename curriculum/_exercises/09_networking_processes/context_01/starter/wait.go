// context_01: Cancellable waits
//
// Wait ignores ctx. Fix it. See README.md.
//
// Reference: https://gobyexample.com/context
package wait

import (
	"context"
	"time"
)

// Wait pauses for d, or until ctx is done, whichever comes first.
func Wait(ctx context.Context, d time.Duration) error {
	time.Sleep(d)
	return nil
}
