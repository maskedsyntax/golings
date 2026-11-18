// string_formatting_01: Lining up columns
//
// Fix the format strings so the output matches README.md exactly.
//
// Reference: https://gobyexample.com/string-formatting
package main

import "fmt"

type line struct {
	item  string
	qty   int
	price float64
}

func main() {
	lines := []line{{"apple", 3, 0.5}, {"banana", 12, 0.25}}

	fmt.Printf("%s|%s|%s\n", "item", "qty", "price")
	for _, l := range lines {
		fmt.Printf("%s|%d|%f\n", l.item, l.qty, l.price)
	}

	fmt.Printf("%s %d %b %d %v\n", "go", 255, 5, 7, 3.5)
}
