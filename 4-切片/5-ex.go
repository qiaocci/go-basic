package main

import "fmt"

func main() {
	b := make([]byte, 512, 512)

	if len(b) == cap(b) {
		b = append(b, 0)[:len(b)]
	}
	fmt.Println(len(b), cap(b))
}
