// Mutant: closes the output as soon as the first input is finished.
package merge

func Merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	done := make(chan struct{}, len(chans))
	for _, c := range chans {
		go func() {
			for v := range c {
				out <- v
			}
			done <- struct{}{}
		}()
	}
	go func() {
		if len(chans) > 0 {
			<-done
		}
		close(out)
	}()
	return out
}
