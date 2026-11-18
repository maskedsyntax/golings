// channel_directions_03: Let the compiler find it
//
// DoubleAll hangs. Add channel directions to worker, then fix what the
// compiler finds. See README.md.
//
// Reference: https://gobyexample.com/channel-directions
package double

func worker(jobs chan int, results chan int) {
	for {
		j := <-jobs
		jobs <- j * 2
	}
}

// DoubleAll returns 2*n for every n, computed by a worker goroutine.
func DoubleAll(nums []int) []int {
	jobs := make(chan int)
	results := make(chan int)
	go worker(jobs, results)
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		jobs <- n
		out = append(out, <-results)
	}
	return out
}
