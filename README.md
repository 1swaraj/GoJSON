# GoJSON
Parsing JSON is a hassle in golang. This package will allow you to parse and search elements in a json without structs.

## Install gojson
go get github.com/swaraj1802/GoJSON/gojson

## Usage

Import
```go
import "github.com/swaraj1802/GoJSON/gojson"
```

Sample Code
```go
jsonParsed, err := gojson.ParseJSON([]byte(`{
   "glossary":{
      "title":"example glossary",
      "GlossDiv":{
         "title":"S",
         "GlossList":{
            "GlossEntry":{
               "ID":"SGML",
               "SortAs":"SGML",
               "GlossTerm":"Standard Generalized Markup Language",
               "Acronym":"SGML",
               "Abbrev":"ISO 8879:1986",
               "GlossDef":{
                  "para":"A meta-markup language, used to create markup languages such as DocBook.",
                  "GlossSeeAlso":[
                     "GML",
                     "XML"
                  ]
               },
               "GlossSee":"markup"
            }
         }
      }
   }
}`))
if err != nil {
   panic(err)
}
value, ok := jsonParsed.Search("glossary", "GlossDiv", "title")
if ok != nil {
  panic(errors.New("Element doesn't exist"))
}
output := value.JSONData()
fmt.Println(output)
```
