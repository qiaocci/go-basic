package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)

	var person = []string{"tom", "jerry"}
	fmt.Println(person)

	// 切片表达式
	numSlice := []int{1, 2, 3, 4, 5}
	s := numSlice[1:3]
	fmt.Println(numSlice, s)
	fmt.Printf("s: %v, len(s): %v, cap(s): %v\n", s, len(s), cap(s))

	// 使用make构造切片
	m := make([]int, 2, 20)
	fmt.Println("===")
	fmt.Println(m)
	fmt.Println(len(m), cap(m))

	var s1 []int
	printSlice2(s1) // len(s1)=0;cap(s1)=0;s1==nil

	s2 := []int{}
	printSlice2(s2)

	s3 := make([]int, 0)
	printSlice2(s3)
	//切片之间是不能比较的，我们不能使用 == 操作符来判断两个切片是否含有全部相等元素。
	//切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	//但是我们不能说一个长度和容量都是0的切片一定是nil

	// 赋值拷贝
	s4 := make([]int, 3)
	s5 := s4
	fmt.Printf("s4: %p, s5: %p\n", s4, s5)
	s4[0] = 1
	fmt.Println(s4, s5) // 下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。

	// append添加元素
	var intSlice []int
	intSlice = append(intSlice, 1)
	printSlice2(intSlice)
	intSlice2 := []int{5, 6, 7}
	intSlice2 = append(intSlice2, intSlice...)
	printSlice2(intSlice2)

	numSlice = []int{}
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	a = []int{30, 31, 32, 33, 34, 35, 36, 37}
	a = append(a[:2], a[3:]...) // 删除第2个元素。 公式： a = append(a[:index], a[index+1:]...)
	fmt.Println(a)
}

func printSlice2(s []int) {
	fmt.Printf("s: %v, len(s): %v, cap(s): %v, s == nil? %v \n", s, len(s), cap(s), s == nil)
}
