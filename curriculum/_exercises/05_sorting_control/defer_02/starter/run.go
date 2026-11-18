// defer_02: When defer looks
//
// The cleanup line is wrong and missing. Fix Run. See README.md.
//
// Reference: https://gobyexample.com/defer
package run

func record(log *[]string, msg string) {
	*log = append(*log, msg)
}

// Run returns a log of "run <step>" lines followed by "cleanup after <last step>".
func Run(steps []string) []string {
	var log []string
	last := ""
	defer record(&log, "cleanup after "+last)
	for _, s := range steps {
		log = append(log, "run "+s)
		last = s
	}
	return log
}
