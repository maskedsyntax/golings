// arrays_01: Arrays are values (reference solution)
package board

// History returns a snapshot of the three-cell board after each move.
func History(moves []int) [][3]int {
	var board [3]int
	var history [][3]int
	for i, m := range moves {
		board[i%3] = m
		history = append(history, board) // appends a copy of the array
	}
	return history
}
