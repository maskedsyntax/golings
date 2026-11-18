package board

import (
	"reflect"
	"testing"
)

func TestHistory(t *testing.T) {
	got := History([]int{7, 8, 9, 1})
	want := [][3]int{{7, 0, 0}, {7, 8, 0}, {7, 8, 9}, {1, 8, 9}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("History([7 8 9 1]) = %v, want %v", got, want)
	}
}

func TestHistorySingle(t *testing.T) {
	if got := History([]int{5}); !reflect.DeepEqual(got, [][3]int{{5, 0, 0}}) {
		t.Errorf("History([5]) = %v, want [[5 0 0]]", got)
	}
}

func TestHistoryEmpty(t *testing.T) {
	if got := History(nil); len(got) != 0 {
		t.Errorf("History(nil) = %v, want no snapshots", got)
	}
}
