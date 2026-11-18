// Mutant: correct result without range (fails the ast check).
package sum

func Sum(ch <-chan int) int {
	total := 0
	for {
		v, ok := <-ch
		if !ok {
			return total
		}
		total += v
	}
}
