// recursion_02: Towers of Hanoi (reference solution)
package hanoi

// Hanoi returns the moves that transfer n disks from peg from to peg to.
func Hanoi(n int, from, to, via string) []string {
	if n == 0 {
		return nil
	}
	moves := Hanoi(n-1, from, via, to)  // clear the way
	moves = append(moves, from+"->"+to) // move the largest disk
	return append(moves, Hanoi(n-1, via, to, from)...)
}
