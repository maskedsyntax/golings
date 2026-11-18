// Mutant: still copies the WaitGroup, papering over it with a sleep (fails vet).
package workers

import (
	"sync"
	"time"
)

func worker(id int, wg sync.WaitGroup, out []int) {
	out[id] = id * id
}

func RunWorkers(n int) []int {
	out := make([]int, n)
	var wg sync.WaitGroup
	for i := range n {
		go worker(i, wg, out)
	}
	time.Sleep(50 * time.Millisecond)
	return out
}
