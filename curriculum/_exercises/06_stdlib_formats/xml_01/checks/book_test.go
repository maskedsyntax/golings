package book

import (
	"encoding/xml"
	"slices"
	"testing"
)

func TestEncode(t *testing.T) {
	got, err := Encode(Book{ID: "b1", Title: "Go", Authors: []string{"Ann", "Bo"}})
	if err != nil {
		t.Fatalf("Encode returned an error: %v", err)
	}
	want := `<book id="b1"><title>Go</title><authors><author>Ann</author><author>Bo</author></authors></book>`
	if got != want {
		t.Fatalf("Encode =\n%s\nwant\n%s", got, want)
	}
}

func TestRoundTrip(t *testing.T) {
	in := `<book id="b2"><title>XML</title><authors><author>Cy</author></authors></book>`
	var b Book
	if err := xml.Unmarshal([]byte(in), &b); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if b.ID != "b2" || b.Title != "XML" || !slices.Equal(b.Authors, []string{"Cy"}) {
		t.Fatalf("Unmarshal gave %+v", b)
	}
}
