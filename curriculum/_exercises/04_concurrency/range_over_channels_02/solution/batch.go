// range_over_channels_02: Batching (reference solution)
package batch

// Batch groups the items from in into slices of size, closing the output when in is closed.
func Batch(in <-chan string, size int) <-chan []string {
	out := make(chan []string)
	go func() {
		defer close(out)
		var b []string
		for s := range in {
			b = append(b, s)
			if len(b) == size {
				out <- b
				b = nil // a fresh slice: the receiver owns the one we sent
			}
		}
		if len(b) > 0 {
			out <- b
		}
	}()
	return out
}
