// channel_buffering_02: The first answer wins (reference solution)
package first

// First calls every source concurrently and returns the first answer.
func First(sources ...func() string) string {
	// Room for every answer: the losers can send and exit even though
	// nobody receives their values.
	ch := make(chan string, len(sources))
	for _, src := range sources {
		go func() {
			ch <- src()
		}()
	}
	return <-ch
}
