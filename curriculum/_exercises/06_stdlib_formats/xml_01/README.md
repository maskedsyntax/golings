# xml_01: Attributes and nesting

`encoding/xml` works like `encoding/json`, and its struct tags have a few extra options:

```go
XMLName xml.Name `xml:"book"`        // name of the element itself
ID      string   `xml:"id,attr"`     // an attribute: <book id="...">
Names   []string `xml:"names>name"`  // <names><name>..</name><name>..</name></names>
```

`Encode` should produce:

```xml
<book id="b1"><title>Go</title><authors><author>Ann</author><author>Bo</author></authors></book>
```

Right now the id is a child element and the authors aren't wrapped in `<authors>`.

## Your task

Fix the struct tags.

Reference: https://gobyexample.com/xml
