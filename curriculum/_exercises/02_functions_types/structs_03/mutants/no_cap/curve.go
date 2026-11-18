// Mutant: modifies in place but forgets the cap at 100.
package curve

type Student struct {
	Name  string
	Score int
}

func Curve(students []Student, bonus int) {
	for i := range students {
		students[i].Score += bonus
	}
}
