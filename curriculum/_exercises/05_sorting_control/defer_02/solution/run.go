// defer_02: When defer looks (reference solution)
package run

// Run returns a log of "run <step>" lines followed by "cleanup after <last step>".
func Run(steps []string) (log []string) {
	last := ""
	defer func() {
		// A closure reads last when it runs, and log is the named result,
		// so this append changes what Run returns.
		log = append(log, "cleanup after "+last)
	}()
	for _, s := range steps {
		log = append(log, "run "+s)
		last = s
	}
	return log
}
