package main

import "fmt"

func main() {
	num0 := 0
	num1 := 1
	num2 := 2
	num3 := 3
	fmt.Printf("num0 p=%p\nnum1 p=%p\nnum2 p=%p\nnum3 p=%p\n", &num0, &num1, &num2, &num3)
	slice := []int{num0, num1, num2, num3}
	myMap := make(map[int]*int)
	for _, value := range slice {
		fmt.Printf("value=%v, p=%p\n", value, &value)
		//num := value
		myMap[value] = &value
	}
	fmt.Println(myMap)
	for k, v := range myMap {
		fmt.Println(k, *v)
	}
}
