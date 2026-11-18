// structs_03: Modifying a copy
//
// Curve has no effect on the students. Fix it.
//
//	s := []Student{{"Ana", 70}, {"Ben", 95}}
//	Curve(s, 10) // s == [{Ana 80} {Ben 100}]
//
// Reference: https://gobyexample.com/structs
package curve

// Student is a name and a test score out of 100.
type Student struct {
	Name  string
	Score int
}

// Curve adds bonus to every student's score in place, capping scores at 100.
func Curve(students []Student, bonus int) {
	for _, s := range students {
		s.Score = min(s.Score+bonus, 100)
	}
}
