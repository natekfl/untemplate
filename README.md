# Untemplate

This is a simple helper to extract values from a string based on a template pattern. Similar to [extract-values](https://github.com/laktek/extract-values)

## Examples

```go
	untemplater1, _ := untemplate.Create("/{year}/{month}/{day}/{title}.html")
	result1, _ := untemplater1.Extract("/2012/08/12/test.html")
    //result1 > map[string]string{"day":"12", "month":"08", "title":"test", "year":"2012"}

	untemplater2, _ := untemplate.Create("{name} <{email}> ({url})")
	result2, _ := untemplater2.Extract("John Doe <john@example.com> (http://example.com)")
	//result2 > map[string]string{"email":"john@example.com", "name":"John Doe", "url":"http://example.com"}

	untemplater3, _ := untemplate.Create("Convert {quantity} {from_unit} to {to_unit}")
	result3, _ := untemplater3.Extract("Convert 1500 Grams to Kilograms")
	//result3 > map[string]string{"from_unit":"Grams", "quantity":"1500", "to_unit":"Kilograms"}
```

## Installation

Install using go cli
```
go get github.com/natekfl/untemplate
```