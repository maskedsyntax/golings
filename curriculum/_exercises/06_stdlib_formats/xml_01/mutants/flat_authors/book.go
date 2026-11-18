// Mutant: attribute fixed, authors not wrapped.
package book

import "encoding/xml"

type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
	Authors []string `xml:"author"`
}

func Encode(b Book) (string, error) {
	out, err := xml.Marshal(b)
	return string(out), err
}
