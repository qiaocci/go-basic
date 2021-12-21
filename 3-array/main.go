package main

import "fmt"

func main() {
	var a [3]int
	var b [4]int

	//a = b //不可以这样做，因为此时a和b是不同的类型
	fmt.Println(a, b)

	// 初始化一
	var numArray = [3]int{1, 2}
	fmt.Println(numArray)

	// 初始化二
	var numArray2 = [...]int{1, 2, 3}
	fmt.Println(numArray2)

	// 初始化三
	var numArray3 = [...]int{1: 1, 3: 5}
	fmt.Println(numArray3)

	for i := 0; i < len(numArray2); i++ {
		fmt.Println(numArray2[i])
	}
	for index, value := range numArray2 {
		fmt.Println(index, value)
	}

	myArray := [3]int{10, 20, 30}
	modifyArray(myArray)
	fmt.Println(myArray) // [10 20 30]
	fmt.Printf("内存地址：%p\n", &myArray) // 内存地址变了。数组的赋值和传参会复制整个数组

}

func modifyArray(x [3]int) {
	fmt.Printf("in function内存地址：%p\n", &x)

	x[0] = 100
	fmt.Println("in modify function")
	fmt.Println(x) // [100 20 30]
}
