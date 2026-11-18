// Mutant: sorts by age only (stable, so ties keep the input order).
package people

import (
	"cmp"
	"slices"
)

type Person struct {
	Name string
	Age  int
}

func SortPeople(people []Person) {
	slices.SortStableFunc(people, func(a, b Person) int { return cmp.Compare(a.Age, b.Age) })
}
