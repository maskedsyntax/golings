// Mutant: no nil case, so nil falls through to "unknown".
package describe

import "fmt"

func Describe(x any) string {
	switch v := x.(type) {
	case int:
		return fmt.Sprintf("int %d", v)
	case string:
		return fmt.Sprintf("string %q", v)
	case []int:
		return fmt.Sprintf("%d ints", len(v))
	case fmt.Stringer:
		return "stringer " + v.String()
	default:
		return "unknown"
	}
}
