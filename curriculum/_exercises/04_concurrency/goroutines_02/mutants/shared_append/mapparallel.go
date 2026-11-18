// Mutant: appends to a shared slice from many goroutines (race, wrong order).
package mapparallel

import "sync"

func MapParallel(items []string, f func(string) string) []string {
	var results []string
	var wg sync.WaitGroup
	for _, item := range items {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results = append(results, f(item))
		}()
	}
	wg.Wait()
	return results
}
