// slices_02: Chunks that don't overlap (reference solution)
package chunk

// Chunk splits items into consecutive slices of at most size elements.
func Chunk(items []int, size int) [][]int {
	if size <= 0 || len(items) == 0 {
		return nil
	}
	var chunks [][]int
	for start := 0; start < len(items); start += size {
		end := min(start+size, len(items))
		// The full slice expression caps capacity so append can't spill over.
		chunks = append(chunks, items[start:end:end])
	}
	return chunks
}
