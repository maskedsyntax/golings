// xml_01: Attributes and nesting
//
// Fix the struct tags so Encode matches README.md.
//
// Reference: https://gobyexample.com/xml
package book

import "encoding/xml"

// Book is a catalog entry.
type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Authors []string `xml:"author"`
}

// Encode returns b as XML.
func Encode(b Book) (string, error) {
	out, err := xml.Marshal(b)
	return string(out), err
}
