// xml_02: Reading a catalog
//
// Add struct tags and implement ParseLibrary. See README.md.
//
// Reference: https://gobyexample.com/xml
package library

import "errors"

// Library is a collection of books.
type Library struct {
	Name  string
	Books []Book
}

// Book is one entry in a library.
type Book struct {
	ID    string
	Title string
	Year  int
}

// ParseLibrary decodes a <library> document.
func ParseLibrary(data []byte) (Library, error) {
	// TODO
	return Library{}, errors.New("not implemented")
}
