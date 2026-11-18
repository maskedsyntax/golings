# xml_02: Reading a catalog

`xml.Unmarshal` fills a struct from a document, using the same tags as `Marshal`:

```xml
<library name="City">
  <book id="1">
    <title>Go Basics</title>
    <meta><year>2021</year></meta>
  </book>
  <book id="2">...</book>
</library>
```

## Your task

Add struct tags to `Library` and `Book`, then implement `ParseLibrary`:

- `Library.Name` comes from the `name` attribute, and `Books` from each `<book>`.
- `Book.ID` comes from the `id` attribute, `Title` from `<title>`, and `Year` from
  `<meta><year>`.
- A document whose root element isn't `<library>` is an error, and so is malformed XML.

Reference: https://gobyexample.com/xml
