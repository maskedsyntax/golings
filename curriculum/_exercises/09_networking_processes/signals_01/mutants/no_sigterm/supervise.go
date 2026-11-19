// Mutant: loops, but only stops on os.Interrupt.
package supervise

import (
	"os"
	"syscall"
)

func Supervise(sigs <-chan os.Signal, reload, shutdown func()) os.Signal {
	for sig := range sigs {
		switch sig {
		case syscall.SIGHUP:
			reload()
		case os.Interrupt:
			shutdown()
			return sig
		}
	}
	return nil
}
