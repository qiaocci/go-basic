package main

import "fmt"

func main() {
	ages := []int{11, 12, 13}
	fmt.Printf("原始slice的内存地址是:%p\n", &ages) // 0xc0000a4018
	fmt.Printf("原始slice的内存地址是:%p\n", ages)  // 0xc0000a4018
	modifySlice(ages)
	fmt.Println("修改后的slice的值是:", ages) // [100 12 13]
}

func modifySlice(ages []int) {
	fmt.Printf("函数内部的slice的内存地址是:%p\n", &ages) // 0xc0000a4030
	ages[0] = 100
}
