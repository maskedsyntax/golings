// time_01: Comparing instants (reference solution)
package booking

import "time"

// Booking reserves the half-open interval [Start, End).
type Booking struct {
	Start, End time.Time
}

// SameMoment reports whether a and b are the same instant.
func SameMoment(a, b time.Time) bool {
	return a.Equal(b) // == would also compare the locations
}

// Overlaps reports whether a and b share any time.
func Overlaps(a, b Booking) bool {
	return a.Start.Before(b.End) && b.Start.Before(a.End)
}
