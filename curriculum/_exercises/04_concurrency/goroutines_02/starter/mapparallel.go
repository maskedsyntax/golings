// goroutines_02: Results from goroutines
//
// MapParallel returns results that are mostly empty. Fix it. See README.md.
//
// Reference: https://gobyexample.com/goroutines
package mapparallel

import "sync"

// MapParallel returns f(items[i]) for every i, computing each in its own goroutine.
func MapParallel(items []string, f func(string) string) []string {
	results := make([]string, len(items))
	var wg sync.WaitGroup
	for i, item := range items {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = f(item)
		}()
	}
	return results
}
