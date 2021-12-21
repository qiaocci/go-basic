package main

import "fmt"

func main() {
	//s := make([]int, 10, 10) // 10 10
	//fmt.Println(len(s), cap(s)) // 11 20
	//
	//s = append(s, 1)
	//fmt.Println(len(s), cap(s))

	b := make([]byte, 10, 10)
	fmt.Println(len(b), cap(b)) // 10 10

	b = append(b, 1)
	fmt.Println(len(b), cap(b)) // 11 24， cap从10扩容到24， 为什么不是两倍？

	c := []byte{1}
	fmt.Println(c)
}
