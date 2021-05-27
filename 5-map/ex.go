package main

import (
	"fmt"
	"strings"
)

func main() {
	ex2()
}

func ex1() {
	s := "how do you do"
	mySlice := strings.Split(s, " ")
	myMap := make(map[string]int, 10)
	for _, word := range mySlice {
		fmt.Println(word, myMap)
		_, ok := myMap[word]
		if !ok {
			myMap[word] = 1
		} else {
			myMap[word] += 1
		}
	}
	fmt.Println(myMap)
}

func ex2() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3) // [1 2 3]
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...) // [1 3]
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
	//切片：包含指针，长度和容量
	//m["q1mi"] 和 s 指向同一个数组：[1 2 3]
	//s 切片删除了下标为 1 的元素，直接在共享数组修改，修改后数组为 [1 3 3]
	//此时 s 的指针，指向底层共享数组，长度为 2，容量为 4（容量为 4，是因为扩容一次）
	//m["q1mi"] 的指针，也指向底层共享数组，长度为 3，容量为 4
}
