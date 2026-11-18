// channels_03: Results in input order
//
// FetchAll returns results in the wrong order when fetches finish at
// different speeds. Fix it. See README.md.
//
// Reference: https://gobyexample.com/channels
package fetch

// FetchAll returns fetch(urls[i]) at index i, fetching concurrently.
func FetchAll(urls []string, fetch func(string) string) []string {
	ch := make(chan string)
	for _, u := range urls {
		go func() {
			ch <- fetch(u)
		}()
	}
	results := make([]string, 0, len(urls))
	for range urls {
		results = append(results, <-ch)
	}
	return results
}
