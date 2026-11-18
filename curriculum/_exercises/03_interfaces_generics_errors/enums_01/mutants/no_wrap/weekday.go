// Mutant: Next doesn't wrap around after Saturday.
package weekday

import "fmt"

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

func (d Weekday) String() string {
	if d < Sunday || d > Saturday {
		return fmt.Sprintf("Weekday(%d)", int(d))
	}
	return names[d]
}

func (d Weekday) IsWeekend() bool { return d == Saturday || d == Sunday }

func (d Weekday) Next() Weekday { return d + 1 }
