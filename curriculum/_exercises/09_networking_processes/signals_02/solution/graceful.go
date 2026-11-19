// signals_02: Shutting down on SIGTERM (reference solution)
package graceful

import (
	"os"
	"os/signal"
	"syscall"
)

// Run calls ready once signals are being handled, waits for SIGINT or
// SIGTERM, then calls cleanup and returns the signal.
func Run(ready, cleanup func()) os.Signal {
	ch := make(chan os.Signal, 1) // Notify never blocks: an unbuffered channel can drop signals
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(ch)
	ready()
	sig := <-ch
	cleanup()
	return sig
}
