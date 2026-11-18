// Mutant: collects in arrival order.
package fetch

func FetchAll(urls []string, fetch func(string) string) []string {
	ch := make(chan string)
	for _, u := range urls {
		go func() { ch <- fetch(u) }()
	}
	var results []string
	for range urls {
		results = append(results, <-ch)
	}
	return results
}
