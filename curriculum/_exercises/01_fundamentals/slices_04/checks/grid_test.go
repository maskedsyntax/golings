package grid

import "testing"

func TestNewGridShape(t *testing.T) {
	g := NewGrid(3, 4)
	if len(g) != 3 {
		t.Fatalf("len(NewGrid(3, 4)) = %d, want 3 rows", len(g))
	}
	for i, row := range g {
		if len(row) != 4 {
			t.Fatalf("row %d has length %d, want 4", i, len(row))
		}
		for j, v := range row {
			if v != 0 {
				t.Errorf("g[%d][%d] = %d, want 0", i, j, v)
			}
		}
	}
}

func TestNewGridRowsAreIndependent(t *testing.T) {
	g := NewGrid(3, 2)
	if len(g) != 3 || len(g[0]) != 2 {
		t.Fatalf("NewGrid(3, 2) has the wrong shape: %v", g)
	}
	g[0][0] = 1
	g[2][1] = 7
	want := [][]int{{1, 0}, {0, 0}, {0, 7}}
	for i := range want {
		for j := range want[i] {
			if g[i][j] != want[i][j] {
				t.Fatalf("after g[0][0] = 1 and g[2][1] = 7, grid is %v, want %v", g, want)
			}
		}
	}
}

func TestNewGridEmpty(t *testing.T) {
	if g := NewGrid(0, 5); len(g) != 0 {
		t.Errorf("NewGrid(0, 5) = %v, want no rows", g)
	}
}
