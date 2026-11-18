// enums_01: iota and String
//
// Implement String, IsWeekend and Next. See README.md.
//
//	Monday.String() == "Monday"
//	Weekday(9).String() == "Weekday(9)"
//	Saturday.Next() == Sunday
//
// Reference: https://gobyexample.com/enums
package weekday

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

// String returns the day's English name.
func (d Weekday) String() string {
	// TODO
	return ""
}

// IsWeekend reports whether d is Saturday or Sunday.
func (d Weekday) IsWeekend() bool {
	// TODO
	return false
}

// Next returns the day after d.
func (d Weekday) Next() Weekday {
	// TODO
	return d
}
