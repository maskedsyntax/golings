// arrays_01: Arrays are values
//
// History records the board after every move, but every snapshot comes out
// identical. Fix it.
//
//	History([]int{7, 8, 9, 1}) == [[7 0 0] [7 8 0] [7 8 9] [1 8 9]]
//
// Reference: https://gobyexample.com/arrays
package board

// History returns a snapshot of the three-cell board after each move.
func History(moves []int) [][3]int {
	board := make([]int, 3)
	var history [][]int
	for i, m := range moves {
		board[i%3] = m
		history = append(history, board)
	}

	out := make([][3]int, len(history))
	for i, snapshot := range history {
		copy(out[i][:], snapshot)
	}
	return out
}
