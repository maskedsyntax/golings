// custom_errors_02: The typed nil error
//
// Check(valid words) != nil, so every caller reports a failure. Fix it.
//
//	Check([]string{"go", "zig"}) == nil
//	Check([]string{"Go"}).Error() == "not lowercase: Go"
//
// Reference: https://gobyexample.com/custom-errors
package check

import (
	"strings"
)

// CheckError lists the words that are not lowercase.
type CheckError struct {
	Problems []string
}

func (e *CheckError) Error() string {
	return "not lowercase: " + strings.Join(e.Problems, ", ")
}

func collect(words []string) *CheckError {
	var e *CheckError
	for _, w := range words {
		if w != strings.ToLower(w) {
			if e == nil {
				e = &CheckError{}
			}
			e.Problems = append(e.Problems, w)
		}
	}
	return e
}

// Check returns nil if every word is lowercase, or a *CheckError otherwise.
func Check(words []string) error {
	return collect(words)
}
