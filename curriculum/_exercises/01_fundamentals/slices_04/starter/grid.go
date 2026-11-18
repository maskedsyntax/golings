// slices_04: The haunted grid
//
// Writing one cell changes every row. Fix NewGrid.
//
// Reference: https://gobyexample.com/slices
package grid

// NewGrid returns a rows x cols grid of zeros.
func NewGrid(rows, cols int) [][]int {
	row := make([]int, cols)
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = row
	}
	return grid
}
