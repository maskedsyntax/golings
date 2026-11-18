// Mutant: n%2 != 1 misclassifies negative odd numbers.
package values

func CanVote(age int, registered bool) bool { return age >= 18 && registered }

func Differ(a, b bool) bool { return a != b }

func IsEven(n int) bool { return n%2 != 1 }
