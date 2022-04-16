package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	s := "abcdefghijklmn"
	s1 := s[:4]
	sHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	s1Header := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Println(sHeader.Len == s1Header.Len)   // false
	fmt.Println(sHeader.Data == s1Header.Data) // true

	s2 := strings.Clone(s[:4])

	s2Header := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Println(sHeader.Len == s2Header.Len)   // false
	fmt.Println(sHeader.Data == s2Header.Data) // false
}
