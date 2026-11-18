// Mutant: the snapshot is taken before the move is applied.
package board

func History(moves []int) [][3]int {
	var board [3]int
	var history [][3]int
	for i, m := range moves {
		history = append(history, board)
		board[i%3] = m
	}
	return history
}
