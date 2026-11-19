// Mutant: catches both signals but skips cleanup.
package graceful

import (
	"os"
	"os/signal"
	"syscall"
)

func Run(ready, cleanup func()) os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(ch)
	ready()
	return <-ch
}
