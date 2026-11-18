// Mutant: starts goroutines but never waits for them.
package runall

func RunAll(tasks ...func()) {
	for _, task := range tasks {
		go task()
	}
}
