// signals_02: Shutting down on SIGTERM
//
// SIGTERM kills the program before cleanup runs. Fix Run. See README.md.
// (Linux and macOS only.)
//
// Reference: https://gobyexample.com/signals
package graceful

import (
	"os"
	"os/signal"
)

// Run calls ready once signals are being handled, waits for SIGINT or
// SIGTERM, then calls cleanup and returns the signal.
func Run(ready, cleanup func()) os.Signal {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	ready()
	sig := <-ch
	cleanup()
	return sig
}
