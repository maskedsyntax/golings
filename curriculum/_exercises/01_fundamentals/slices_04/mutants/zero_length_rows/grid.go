// Mutant: rows have capacity but no length, so indexing panics.
package grid

func NewGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, 0, cols)
	}
	return grid
}
