// non_blocking_channel_operations_01: Try, don't wait
//
// Implement TrySend and TryReceive. See README.md.
//
// Reference: https://gobyexample.com/non-blocking-channel-operations
package try

// TrySend sends msg on ch if that can happen immediately.
func TrySend(ch chan<- string, msg string) bool {
	// TODO
	return false
}

// TryReceive receives from ch if a value is ready immediately.
func TryReceive(ch <-chan string) (string, bool) {
	// TODO
	return "", false
}
