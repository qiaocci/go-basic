package main

import "fmt"
import (
	"github.com/Jeffail/gabs/v2"
)

func main() {
	jsonParsed, err := gabs.ParseJSON([]byte(`{
			"outter":{
				"inner":{
					"value1":10,
					"value2":22
				},
				"alsoInner":{
					"value1":20,
					"array1":[
						30, 40
					]
				}
			}
		}`))
	if err != nil {
		panic(err)
	}

	var value float64
	var ok bool

	value, ok = jsonParsed.Path("outter.inner.value1").Data().(float64)
	fmt.Printf("value=%v, type=%T, ok=%v\n", value, value, ok)

	value, ok = jsonParsed.Search("outter", "inner", "value1").Data().(float64)
	fmt.Printf("value=%v, type=%T\n", value, value)

	value, ok = jsonParsed.Search("outter", "alsoInner", "array1", "1").Data().(float64)
	fmt.Printf("value=%v, type=%T\n", value, value)

	gObj, err := jsonParsed.JSONPointer("/outter/alsoInner/array1/1")
	if err != nil {
		panic(err)
	}
	value, ok = gObj.Data().(float64)
	fmt.Printf("value=%v, ok=%v\n", value, ok)

	exists := jsonParsed.Exists("outter", "inner", "value1")
	fmt.Println(exists)
}
