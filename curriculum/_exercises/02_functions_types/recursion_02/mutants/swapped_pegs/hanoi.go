// Mutant: the second recursive call uses the wrong spare peg.
package hanoi

func Hanoi(n int, from, to, via string) []string {
	if n == 0 {
		return nil
	}
	moves := Hanoi(n-1, from, via, to)
	moves = append(moves, from+"->"+to)
	return append(moves, Hanoi(n-1, via, from, to)...)
}
