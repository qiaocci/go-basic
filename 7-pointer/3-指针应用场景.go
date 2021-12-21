package main

import "fmt"

func main() {
	ex2()
}

func ex1() {
	v1 := "武沛齐"
	v2 := v1 // 内存中拷贝了一份
	v1 = "alex"
	fmt.Println(v1, v2) // alex 武沛齐

}

// 如果希望数据修改后，全部同步，用指针
// 反之就不用
func ex2() {
	v1 := "武沛齐"
	v2 := &v1
	v1 = "alex"
	fmt.Println(v1, *v2) // alex alex
}
