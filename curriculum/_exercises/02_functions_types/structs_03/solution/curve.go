// structs_03: Modifying a copy (reference solution)
package curve

// Student is a name and a test score out of 100.
type Student struct {
	Name  string
	Score int
}

// Curve adds bonus to every student's score in place, capping scores at 100.
func Curve(students []Student, bonus int) {
	for i := range students {
		s := &students[i] // a pointer to the element, not a copy
		s.Score = min(s.Score+bonus, 100)
	}
}
