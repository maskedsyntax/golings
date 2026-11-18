// Mutant: authors wrapped, id still an element.
package book

import "encoding/xml"

type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Authors []string `xml:"authors>author"`
}

func Encode(b Book) (string, error) {
	out, err := xml.Marshal(b)
	return string(out), err
}
