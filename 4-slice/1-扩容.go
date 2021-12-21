package main

import "fmt"

func main() {
	b := make([]byte, 0, 10)
	printSlice3(b)
	for i := 0; i < 10; i++ {
		b = append(b, 0)
	}
	fmt.Println("最终")
	printSlice3(b)

	b = append(b, 1) // s: [0 0 0 0 0 0 0 0 0 0 1], len(s): 11, cap(s): 24, s == nil? false
	printSlice3(b)
	//printSlice3(append(b, 1)[:len(b)]) //s: [0 0 0 0 0 0 0 0 0 0], len(s): 10, cap(s): 24, s == nil? false

	fmt.Println("======")
	slice1 := make([]int, 1)
	fmt.Println("cap of slice1", cap(slice1))
	slice1 = append(slice1, 1)
	fmt.Println("cap of slice1", cap(slice1))
	slice1 = append(slice1, 2)
	fmt.Println("cap of slice1", cap(slice1))

	fmt.Println("======")
	slice10 := make([]int, 10)
	fmt.Println("cap of slice10", cap(slice10))
	slice10 = append(slice10, 1)
	fmt.Println("cap of slice10", cap(slice10))
	//slice1 = append(slice1,2)
	//fmt.Println("cap of slice1",cap(slice1))

	fmt.Println()

	slice1024 := make([]int, 1024)
	fmt.Println("cap of slice1024", cap(slice1024))
	slice1024 = append(slice1024, 1)
	fmt.Println("cap of slice1024", cap(slice1024))
	slice1024 = append(slice1024, 2)
	fmt.Println("cap of slice1024", cap(slice1024))

}
func printSlice3(s []byte) {
	fmt.Printf("s: %v, len(s): %v, cap(s): %v, s == nil? %v \n", s, len(s), cap(s), s == nil)
}
