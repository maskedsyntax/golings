package library

import "testing"

const doc = `<?xml version="1.0" encoding="UTF-8"?>
<library name="City">
  <book id="1" lang="en">
    <title>Go Basics</title>
    <meta><year>2021</year><pages>200</pages></meta>
  </book>
  <book id="2">
    <title>XML Deep Dive</title>
    <meta><year>2019</year></meta>
  </book>
</library>
`

func TestParseLibrary(t *testing.T) {
	lib, err := ParseLibrary([]byte(doc))
	if err != nil {
		t.Fatalf("ParseLibrary returned an error: %v", err)
	}
	if lib.Name != "City" {
		t.Errorf("Name = %q, want \"City\"", lib.Name)
	}
	want := []Book{{"1", "Go Basics", 2021}, {"2", "XML Deep Dive", 2019}}
	if len(lib.Books) != len(want) {
		t.Fatalf("got %d books, want %d: %+v", len(lib.Books), len(want), lib.Books)
	}
	for i, b := range lib.Books {
		if b != want[i] {
			t.Errorf("book %d = %+v, want %+v", i, b, want[i])
		}
	}
}

func TestParseLibraryErrors(t *testing.T) {
	for _, in := range []string{
		`<shop name="x"><book id="1"><title>T</title></book></shop>`,
		`<library name="x"><book id="1">`,
		``,
	} {
		if _, err := ParseLibrary([]byte(in)); err == nil {
			t.Errorf("ParseLibrary(%q) returned no error", in)
		}
	}
}
