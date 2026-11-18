// enums_02: Names out of sync (reference solution)
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

// Keyed by the constants themselves, so inserting a new constant can't shift names.
var names = map[Status]string{
	Pending:   "pending",
	Active:    "active",
	Suspended: "suspended",
	Closed:    "closed",
}

// String returns the status name.
func (s Status) String() string {
	if name, ok := names[s]; ok {
		return name
	}
	return fmt.Sprintf("Status(%d)", int(s))
}

// ParseStatus returns the Status with the given name.
func ParseStatus(name string) (Status, bool) {
	for s, n := range names {
		if n == name {
			return s, true
		}
	}
	return 0, false
}
