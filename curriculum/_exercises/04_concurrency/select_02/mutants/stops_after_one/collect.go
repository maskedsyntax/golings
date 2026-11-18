// Mutant: returns after the first value.
package collect

func Collect(data <-chan int, quit <-chan struct{}) []int {
	select {
	case v := <-data:
		return []int{v}
	case <-quit:
		return nil
	}
}
