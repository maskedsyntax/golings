// Mutant: always returns a *CheckError, empty when everything is fine.
package check

import "strings"

type CheckError struct {
	Problems []string
}

func (e *CheckError) Error() string { return "not lowercase: " + strings.Join(e.Problems, ", ") }

func Check(words []string) error {
	e := &CheckError{}
	for _, w := range words {
		if w != strings.ToLower(w) {
			e.Problems = append(e.Problems, w)
		}
	}
	return e
}
