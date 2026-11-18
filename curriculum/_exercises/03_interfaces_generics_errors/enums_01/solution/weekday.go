// enums_01: iota and String (reference solution)
package weekday

import "fmt"

// Weekday is a day of the week.
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var names = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

// String returns the day's English name.
func (d Weekday) String() string {
	if d < Sunday || d > Saturday {
		return fmt.Sprintf("Weekday(%d)", int(d))
	}
	return names[d]
}

// IsWeekend reports whether d is Saturday or Sunday.
func (d Weekday) IsWeekend() bool {
	return d == Saturday || d == Sunday
}

// Next returns the day after d.
func (d Weekday) Next() Weekday {
	return (d + 1) % 7
}
