// Mutant: ParseStatus accepts any capitalization.
package status

import (
	"fmt"
	"strings"
)

type Status int

const (
	Pending Status = iota
	Active
	Suspended
	Closed
)

var names = map[Status]string{Pending: "pending", Active: "active", Suspended: "suspended", Closed: "closed"}

func (s Status) String() string {
	if name, ok := names[s]; ok {
		return name
	}
	return fmt.Sprintf("Status(%d)", int(s))
}

func ParseStatus(name string) (Status, bool) {
	for s, n := range names {
		if strings.EqualFold(n, name) {
			return s, true
		}
	}
	return 0, false
}
