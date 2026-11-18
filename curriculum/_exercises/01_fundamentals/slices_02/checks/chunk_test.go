package chunk

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		size  int
		want  [][]int
	}{
		{"uneven", []int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{"even", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{"size larger than input", []int{1, 2}, 5, [][]int{{1, 2}}},
		{"size one", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{"zero size", []int{1, 2}, 0, nil},
		{"negative size", []int{1, 2}, -1, nil},
		{"empty", nil, 3, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.items, tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk(%v, %d) = %v, want %v", tt.items, tt.size, got, tt.want)
			}
		})
	}
}

func TestChunkAppendDoesNotOverwrite(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	chunks := Chunk(items, 2)
	_ = append(chunks[0], 99)
	if !reflect.DeepEqual(items, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("appending to the first chunk changed items to %v", items)
	}
	if !reflect.DeepEqual(chunks[1], []int{3, 4}) {
		t.Errorf("appending to the first chunk changed the second chunk to %v", chunks[1])
	}
}
