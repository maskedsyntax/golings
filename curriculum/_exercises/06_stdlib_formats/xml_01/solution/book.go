// xml_01: Attributes and nesting (reference solution)
package book

import "encoding/xml"

// Book is a catalog entry.
type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
	Authors []string `xml:"authors>author"`
}

// Encode returns b as XML.
func Encode(b Book) (string, error) {
	out, err := xml.Marshal(b)
	return string(out), err
}
