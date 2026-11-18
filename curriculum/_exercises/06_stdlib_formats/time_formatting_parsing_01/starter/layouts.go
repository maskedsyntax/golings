// time_formatting_parsing_01: Reference-time layouts
//
// Fix the three layouts. See README.md.
//
// Reference: https://gobyexample.com/time-formatting-parsing
package layouts

import "time"

// FormatDate formats t as year-month-day, e.g. "2024-03-05".
func FormatDate(t time.Time) string {
	return t.Format("YYYY-MM-DD")
}

// ParseDate parses a year-month-day date such as "2024-03-05".
func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04", s)
}

// Clock formats the time of day on a 12-hour clock, e.g. "3:04 PM".
func Clock(t time.Time) string {
	return t.Format("15:04 PM")
}
