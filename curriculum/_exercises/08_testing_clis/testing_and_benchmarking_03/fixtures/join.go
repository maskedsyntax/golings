package join

import "strings"

// These are declared as variables so golings can count how often your
// benchmark calls them. Call them like any other function.

// JoinPlus concatenates parts with +=, copying the result every time.
var JoinPlus = func(parts []string) string {
	s := ""
	for _, p := range parts {
		s += p
	}
	return s
}

// JoinBuilder concatenates parts with a strings.Builder.
var JoinBuilder = func(parts []string) string {
	var b strings.Builder
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String()
}
