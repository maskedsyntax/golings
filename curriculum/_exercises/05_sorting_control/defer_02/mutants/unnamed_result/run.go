// Mutant: a closure, but the result isn't named, so the append is lost.
package run

func Run(steps []string) []string {
	var log []string
	last := ""
	defer func() {
		log = append(log, "cleanup after "+last)
	}()
	for _, s := range steps {
		log = append(log, "run "+s)
		last = s
	}
	return log
}
