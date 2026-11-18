// Mutant: keeps the values keep rejects.
package pipeline

func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func Filter(in <-chan int, keep func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if !keep(v) {
				out <- v
			}
		}
	}()
	return out
}

func Count(in <-chan int) int {
	n := 0
	for range in {
		n++
	}
	return n
}
