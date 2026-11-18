// Mutant: reuses the batch's backing array, corrupting batches already sent.
package batch

func Batch(in <-chan string, size int) <-chan []string {
	out := make(chan []string, 10)
	go func() {
		defer close(out)
		b := make([]string, 0, size)
		for s := range in {
			b = append(b, s)
			if len(b) == size {
				out <- b
				b = b[:0]
			}
		}
		if len(b) > 0 {
			out <- b
		}
	}()
	return out
}
