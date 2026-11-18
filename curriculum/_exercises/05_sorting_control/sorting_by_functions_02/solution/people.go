// sorting_by_functions_02: Sorting records (reference solution)
package people

import (
	"cmp"
	"slices"
	"strings"
)

// Person is a name and an age.
type Person struct {
	Name string
	Age  int
}

// SortPeople sorts people in place by age, then by name.
func SortPeople(people []Person) {
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Or(cmp.Compare(a.Age, b.Age), strings.Compare(a.Name, b.Name))
	})
}
