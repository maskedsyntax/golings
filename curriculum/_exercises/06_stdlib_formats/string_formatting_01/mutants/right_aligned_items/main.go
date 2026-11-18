// Mutant: item column right-aligned.
package main

import "fmt"

func main() {
	fmt.Printf("%10s|%5s|%8s\n", "item", "qty", "price")
	fmt.Printf("%10s|%5d|%8.2f\n", "apple", 3, 0.5)
	fmt.Printf("%10s|%5d|%8.2f\n", "banana", 12, 0.25)
	fmt.Printf("%q %x %08b %+d %T\n", "go", 255, 5, 7, 3.5)
}
