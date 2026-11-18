// Mutant: TrySend reports success without checking.
package try

func TrySend(ch chan<- string, msg string) bool {
	select {
	case ch <- msg:
	default:
	}
	return true
}

func TryReceive(ch <-chan string) (string, bool) {
	select {
	case m := <-ch:
		return m, true
	default:
		return "", false
	}
}
