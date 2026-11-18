// Mutant: forgets the final partial batch.
package batch

func Batch(in <-chan string, size int) <-chan []string {
	out := make(chan []string)
	go func() {
		defer close(out)
		var b []string
		for s := range in {
			b = append(b, s)
			if len(b) == size {
				out <- b
				b = nil
			}
		}
	}()
	return out
}
