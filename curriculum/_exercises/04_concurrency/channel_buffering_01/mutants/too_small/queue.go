// Mutant: buffer one short; blocks only on the final send (test has 3 msgs).
package queue

func Queue(msgs ...string) chan string {
	ch := make(chan string, max(0, len(msgs)-1))
	for _, m := range msgs {
		ch <- m
	}
	return ch
}
