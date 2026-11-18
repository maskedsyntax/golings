// non_blocking_channel_operations_01: Try, don't wait (reference solution)
package try

// TrySend sends msg on ch if that can happen immediately.
func TrySend(ch chan<- string, msg string) bool {
	select {
	case ch <- msg:
		return true
	default:
		return false
	}
}

// TryReceive receives from ch if a value is ready immediately.
func TryReceive(ch <-chan string) (string, bool) {
	select {
	case m := <-ch:
		return m, true
	default:
		return "", false
	}
}
