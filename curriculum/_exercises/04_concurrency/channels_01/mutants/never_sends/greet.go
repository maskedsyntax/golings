// Mutant: returns a channel nobody sends on.
package greet

func Greet(name string) <-chan string {
	return make(chan string)
}
