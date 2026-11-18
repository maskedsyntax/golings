package hanoi

import (
	"slices"
	"strings"
	"testing"
)

func TestHanoiSmall(t *testing.T) {
	if got := Hanoi(0, "A", "C", "B"); len(got) != 0 {
		t.Errorf("Hanoi(0) = %v, want no moves", got)
	}
	if got := Hanoi(1, "A", "C", "B"); !slices.Equal(got, []string{"A->C"}) {
		t.Errorf("Hanoi(1, A, C, B) = %v, want [A->C]", got)
	}
	if got := Hanoi(2, "A", "C", "B"); !slices.Equal(got, []string{"A->B", "A->C", "B->C"}) {
		t.Errorf("Hanoi(2, A, C, B) = %v, want [A->B A->C B->C]", got)
	}
}

// replay checks every move against the rules and returns the final pegs.
func replay(t *testing.T, n int, moves []string) map[string][]int {
	t.Helper()
	pegs := map[string][]int{"L": nil, "M": nil, "R": nil}
	for d := n; d >= 1; d-- {
		pegs["L"] = append(pegs["L"], d)
	}
	for i, m := range moves {
		from, to, ok := strings.Cut(m, "->")
		_, knownFrom := pegs[from]
		_, knownTo := pegs[to]
		if !ok || !knownFrom || !knownTo || from == to {
			t.Fatalf("move %d is %q, want the form \"from->to\" with pegs L, M, R", i+1, m)
		}
		src := pegs[from]
		if len(src) == 0 {
			t.Fatalf("move %d (%s) takes a disk from an empty peg", i+1, m)
		}
		disk := src[len(src)-1]
		if dst := pegs[to]; len(dst) > 0 && dst[len(dst)-1] < disk {
			t.Fatalf("move %d (%s) puts disk %d on smaller disk %d", i+1, m, disk, dst[len(dst)-1])
		}
		pegs[from] = src[:len(src)-1]
		pegs[to] = append(pegs[to], disk)
	}
	return pegs
}

func TestHanoiIsLegalAndOptimal(t *testing.T) {
	for _, n := range []int{3, 5, 8} {
		moves := Hanoi(n, "L", "R", "M")
		if want := 1<<n - 1; len(moves) != want {
			t.Errorf("Hanoi(%d) made %d moves, want %d", n, len(moves), want)
		}
		pegs := replay(t, n, moves)
		if len(pegs["R"]) != n {
			t.Errorf("after Hanoi(%d), peg R holds %v, want all %d disks", n, pegs["R"], n)
		}
	}
}
