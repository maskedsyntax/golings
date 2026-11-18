// waitgroups_02: Don't copy the WaitGroup (reference solution)
package workers

import "sync"

func worker(id int, wg *sync.WaitGroup, out []int) {
	defer wg.Done()
	out[id] = id * id
}

// RunWorkers starts n workers and returns their results once all are done.
func RunWorkers(n int) []int {
	out := make([]int, n)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go worker(i, &wg, out)
	}
	wg.Wait()
	return out
}
