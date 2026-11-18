// Mutant: SameMoment fixed, touching bookings still overlap.
package booking

import "time"

type Booking struct {
	Start, End time.Time
}

func SameMoment(a, b time.Time) bool {
	return a.Equal(b)
}

func Overlaps(a, b Booking) bool {
	return !a.End.Before(b.Start) && !b.End.Before(a.Start)
}
