// Mutant: accepts any root element.
package library

import "encoding/xml"

type Library struct {
	Name  string `xml:"name,attr"`
	Books []Book `xml:"book"`
}

type Book struct {
	ID    string `xml:"id,attr"`
	Title string `xml:"title"`
	Year  int    `xml:"meta>year"`
}

func ParseLibrary(data []byte) (Library, error) {
	var lib Library
	err := xml.Unmarshal(data, &lib)
	return lib, err
}
