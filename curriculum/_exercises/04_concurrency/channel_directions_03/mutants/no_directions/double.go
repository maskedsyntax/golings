// Mutant: bug fixed but parameters left bidirectional.
package double

func worker(jobs chan int, results chan int) {
	for {
		j := <-jobs
		results <- j * 2
	}
}

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
