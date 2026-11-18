// time_01: Comparing instants
//
// Fix SameMoment and Overlaps. See README.md.
//
// Reference: https://gobyexample.com/time
package booking

import "time"

// Booking reserves the half-open interval [Start, End).
type Booking struct {
	Start, End time.Time
}

// SameMoment reports whether a and b are the same instant.
func SameMoment(a, b time.Time) bool {
	return a == b
}

// Overlaps reports whether a and b share any time.
func Overlaps(a, b Booking) bool {
	return !a.End.Before(b.Start) && !b.End.Before(a.Start)
}
