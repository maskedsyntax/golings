// Mutant: buffered, but SIGTERM still kills the program.
package graceful

import (
	"os"
	"os/signal"
)

func Run(ready, cleanup func()) os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	defer signal.Stop(ch)
	ready()
	sig := <-ch
	cleanup()
	return sig
}
