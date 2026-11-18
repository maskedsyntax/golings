package sum

import "testing"

func TestParallelSum(t *testing.T) {
	chunks := make([][]int, 50)
	want := 0
	for i := range chunks {
		for j := range 200 {
			chunks[i] = append(chunks[i], i+j)
			want += i + j
		}
	}
	for range 5 {
		if got := ParallelSum(chunks); got != want {
			t.Fatalf("ParallelSum = %d, want %d (updates were lost)", got, want)
		}
	}
}

func TestParallelSumSmall(t *testing.T) {
	if got := ParallelSum([][]int{{1, 2}, {3}, {}}); got != 6 {
		t.Errorf("ParallelSum({{1,2},{3},{}}) = %d, want 6", got)
	}
	if got := ParallelSum(nil); got != 0 {
		t.Errorf("ParallelSum(nil) = %d, want 0", got)
	}
}
