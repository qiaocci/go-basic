package main

import "fmt"

func main() {

	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]
	fmt.Printf("Slice_a: %v, Array_a: %v\n", Slice_a, Array_a)

	Slice_a[0] = 'z'
	fmt.Printf("Slice_a: %v, Array_a: %v\n", Slice_a, Array_a)

	for i := 0; i < 10; i++ {
		Slice_a = append(Slice_a, byte(i))
	}
	fmt.Printf("Slice_a: %v, Array_a: %v\n", Slice_a, Array_a)

}
