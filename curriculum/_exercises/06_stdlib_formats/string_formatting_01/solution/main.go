// string_formatting_01: Lining up columns (reference solution)
package main

import "fmt"

type line struct {
	item  string
	qty   int
	price float64
}

func main() {
	lines := []line{{"apple", 3, 0.5}, {"banana", 12, 0.25}}

	fmt.Printf("%-10s|%5s|%8s\n", "item", "qty", "price")
	for _, l := range lines {
		fmt.Printf("%-10s|%5d|%8.2f\n", l.item, l.qty, l.price)
	}

	fmt.Printf("%q %x %08b %+d %T\n", "go", 255, 5, 7, 3.5)
}
