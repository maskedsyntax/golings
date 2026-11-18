// Mutant: runs the steps directly.
package relay

func Relay(steps ...func()) {
	for _, step := range steps {
		step()
	}
}
