// Mutant: closes quit but returns without waiting for the workers.
package workers

func Start(n int, work func(id int)) (stop func()) {
	quit := make(chan struct{})
	for id := range n {
		go func() {
			for {
				select {
				case <-quit:
					return
				default:
					work(id)
				}
			}
		}()
	}
	return func() { close(quit) }
}
