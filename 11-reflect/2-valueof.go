package main

import (
	"fmt"
	"reflect"
)

func reflectValueOf(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	default:
		fmt.Printf("类型未找到")
	}
}
func main() {
	var a float32 = 3.14
	var b int64 = 100

	reflectValueOf(a)
	reflectValueOf(b)

	c := reflect.ValueOf(10)
	fmt.Printf("type: %T\n", c.Kind())
}
