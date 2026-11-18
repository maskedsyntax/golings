// xml_02: Reading a catalog (reference solution)
package library

import "encoding/xml"

// Library is a collection of books.
type Library struct {
	XMLName xml.Name `xml:"library"` // rejects any other root element
	Name    string   `xml:"name,attr"`
	Books   []Book   `xml:"book"`
}

// Book is one entry in a library.
type Book struct {
	ID    string `xml:"id,attr"`
	Title string `xml:"title"`
	Year  int    `xml:"meta>year"`
}

// ParseLibrary decodes a <library> document.
func ParseLibrary(data []byte) (Library, error) {
	var lib Library
	if err := xml.Unmarshal(data, &lib); err != nil {
		return Library{}, err
	}
	return lib, nil
}
