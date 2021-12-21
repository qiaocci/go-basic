package main

import "fmt"

func main() {
	ex2()
}

func ex2() {
	fmt.Println("aaaa")
	intArray := [...]int{1, 3, 5, 7, 8}
	for index1, v1 := range intArray {
		for index2, v2 := range intArray {
			if v1+v2 == 8 && index1 <= index2 {
				fmt.Println(index1, index2)
			}
		}
	}
}

func ex1() {
	myArray := [...]int{1, 3, 5, 7, 9}
	total := 0
	for _, value := range myArray {
		total += value
	}
	fmt.Println(total)
}
