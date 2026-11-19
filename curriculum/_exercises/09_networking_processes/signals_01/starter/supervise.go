// signals_01: Reacting to signals
//
// Supervise stops after one signal and ignores SIGTERM. Fix it. See README.md.
//
// Reference: https://gobyexample.com/signals
package supervise

import (
	"os"
	"syscall"
)

// Supervise handles signals until it's told to stop, and returns the stopping signal.
func Supervise(sigs <-chan os.Signal, reload, shutdown func()) os.Signal {
	sig := <-sigs
	switch sig {
	case syscall.SIGHUP:
		reload()
	case os.Interrupt:
		shutdown()
	}
	return sig
}
