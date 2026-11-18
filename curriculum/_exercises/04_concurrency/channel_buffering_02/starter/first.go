// channel_buffering_02: The first answer wins
//
// First leaks a goroutine for every source that loses. Fix it. See README.md.
//
// Reference: https://gobyexample.com/channel-buffering
package first

// First calls every source concurrently and returns the first answer.
func First(sources ...func() string) string {
	ch := make(chan string)
	for _, src := range sources {
		go func() {
			ch <- src()
		}()
	}
	return <-ch
}
