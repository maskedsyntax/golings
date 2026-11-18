// Mutant: starts every step at once and waits for all of them.
package relay

import "sync"

func Relay(steps ...func()) {
	var wg sync.WaitGroup
	for _, step := range steps {
		wg.Add(1)
		go func() { defer wg.Done(); step() }()
	}
	wg.Wait()
}
