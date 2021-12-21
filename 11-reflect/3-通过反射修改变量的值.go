package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 { // 反射中使用专有的Elem()方法来获取指针对应的值
		v.Elem().SetInt(200)
	}
}
func main() {
	var a int64 = 100

	reflectSetValue2(a)
	fmt.Println(a)
}
