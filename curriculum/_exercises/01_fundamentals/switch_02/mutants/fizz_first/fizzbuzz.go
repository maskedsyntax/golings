// Mutant: checking 3 before 15 means FizzBuzz is never produced.
package fizzbuzz

import "strconv"

func FizzBuzz(n int) string {
	switch {
	case n%3 == 0:
		return "Fizz"
	case n%15 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
