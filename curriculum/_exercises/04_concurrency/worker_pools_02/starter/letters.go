// worker_pools_02: Waiting in the wrong place
//
// CountLetters hangs. Fix it. See README.md.
//
// Reference: https://gobyexample.com/worker-pools
package letters

import "sync"

type result struct {
	word string
	n    int
}

// CountLetters returns measure(w) for every word, using the given number of workers.
func CountLetters(words []string, workers int, measure func(string) int) map[string]int {
	jobs := make(chan string)
	results := make(chan result)
	go func() {
		defer close(jobs)
		for _, w := range words {
			jobs <- w
		}
	}()
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range jobs {
				results <- result{w, measure(w)}
			}
		}()
	}
	wg.Wait()
	close(results)
	out := map[string]int{}
	for r := range results {
		out[r.word] = r.n
	}
	return out
}
