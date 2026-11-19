// signals_01: Reacting to signals (reference solution)
package supervise

import (
	"os"
	"syscall"
)

// Supervise handles signals until it's told to stop, and returns the stopping signal.
func Supervise(sigs <-chan os.Signal, reload, shutdown func()) os.Signal {
	for sig := range sigs { // ends if the channel is closed
		switch sig {
		case syscall.SIGHUP:
			reload()
		case os.Interrupt, syscall.SIGTERM:
			shutdown()
			return sig
		}
	}
	return nil
}
