package main

import (
	"fmt"
	"strings"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()
	fmt.Println(nextNumber)

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	fmt.Println(strings.Repeat("=", 20))
	nextNumber2 := getSequence()
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())

}
