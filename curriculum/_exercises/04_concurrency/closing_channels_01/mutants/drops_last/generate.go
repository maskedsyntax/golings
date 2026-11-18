// Mutant: closes one value too early.
package generate

func Generate(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i, n := range nums {
			if i == len(nums)-1 {
				return
			}
			ch <- n
		}
	}()
	return ch
}
