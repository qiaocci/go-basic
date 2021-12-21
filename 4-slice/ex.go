package main

import (
	"fmt"
	"sort"
)
func main() {
	ex2()
}

func ex1() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}

func ex2()  {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}