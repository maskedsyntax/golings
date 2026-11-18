// Mutant: closes before the task runs.
package start

func Start(task func()) <-chan struct{} {
	done := make(chan struct{})
	close(done)
	go task()
	return done
}
