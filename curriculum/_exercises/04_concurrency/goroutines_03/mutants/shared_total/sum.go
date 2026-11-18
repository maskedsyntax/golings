// Mutant: every goroutine updates the same total.
package sum

import "sync"

func ParallelSum(chunks [][]int) int {
	total := 0
	var wg sync.WaitGroup
	for _, chunk := range chunks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := 0
			for _, n := range chunk {
				s += n
			}
			total += s
		}()
	}
	wg.Wait()
	return total
}
