// slices_04: The haunted grid (reference solution)
package grid

// NewGrid returns a rows x cols grid of zeros.
func NewGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	return grid
}
