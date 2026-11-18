// Mutant: Generate sends in reverse order.
package pipe

func Generate(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := len(nums) - 1; i >= 0; i-- {
			ch <- nums[i]
		}
	}()
	return ch
}

func Take(ch <-chan int, n int) []int {
	out := make([]int, 0, n)
	for range n {
		out = append(out, <-ch)
	}
	return out
}
