package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}

// Float64bits 类型转换
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func main() {
	f := 3.14 // float64
	fmt.Printf("f=%v, type=%T\n", f, f)
	u := Float64bits(f)
	fmt.Printf("u=%v, type=%T\n", u, u)

}
