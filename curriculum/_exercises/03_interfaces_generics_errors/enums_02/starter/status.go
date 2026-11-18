// enums_02: Names out of sync
//
// Suspended prints as "closed" since it was added. Fix it.
//
//	Suspended.String() == "suspended"
//	ParseStatus("closed") == Closed, true
//
// Reference: https://gobyexample.com/enums
package status

import "fmt"

// Status is the state of an account.
type Status int

const (
	Pending Status = iota
	Active
	Suspended // added in v2
	Closed
)

var names = []string{"pending", "active", "closed"}

// String returns the status name.
func (s Status) String() string {
	if s >= 0 && int(s) < len(names) {
		return names[s]
	}
	return fmt.Sprintf("Status(%d)", int(s))
}

// ParseStatus returns the Status with the given name.
func ParseStatus(name string) (Status, bool) {
	for i, n := range names {
		if n == name {
			return Status(i), true
		}
	}
	return 0, false
}
