// Mutant: Overlaps fixed, SameMoment still uses ==.
package booking

import "time"

type Booking struct {
	Start, End time.Time
}

func SameMoment(a, b time.Time) bool {
	return a == b
}

func Overlaps(a, b Booking) bool {
	return a.Start.Before(b.End) && b.Start.Before(a.End)
}
