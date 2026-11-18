// Mutant: chunks keep the capacity of items, so append overwrites neighbors.
package chunk

func Chunk(items []int, size int) [][]int {
	if size <= 0 || len(items) == 0 {
		return nil
	}
	var chunks [][]int
	for start := 0; start < len(items); start += size {
		end := min(start+size, len(items))
		chunks = append(chunks, items[start:end])
	}
	return chunks
}
