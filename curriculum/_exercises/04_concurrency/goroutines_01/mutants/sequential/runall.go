// Mutant: runs the tasks one at a time.
package runall

func RunAll(tasks ...func()) {
	for _, task := range tasks {
		task()
	}
}
