// channels_03: Results in input order (reference solution)
package fetch

type result struct {
	i    int
	body string
}

// FetchAll returns fetch(urls[i]) at index i, fetching concurrently.
func FetchAll(urls []string, fetch func(string) string) []string {
	ch := make(chan result)
	for i, u := range urls {
		go func() {
			ch <- result{i, fetch(u)}
		}()
	}
	results := make([]string, len(urls))
	for range urls {
		r := <-ch
		results[r.i] = r.body
	}
	return results
}
