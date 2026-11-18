package people

import (
	"slices"
	"testing"
)

func TestSortPeople(t *testing.T) {
	people := []Person{
		{"Zoe", 30}, {"Alex", 72}, {"Mia", 25}, {"Ben", 30}, {"Liam", 25}, {"Ada", 30},
	}
	SortPeople(people)
	want := []Person{
		{"Liam", 25}, {"Mia", 25}, {"Ada", 30}, {"Ben", 30}, {"Zoe", 30}, {"Alex", 72},
	}
	if !slices.Equal(people, want) {
		t.Fatalf("SortPeople =\n  %v\nwant\n  %v", people, want)
	}
}
