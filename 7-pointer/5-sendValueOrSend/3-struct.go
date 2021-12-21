package main

import "fmt"

func main() {
	p := Person{Name: "qiaocc"}

	fmt.Printf("原始struct的内存地址是:%p\n", &p)
	modifyStruct(&p)
	fmt.Println("修改后的struct的值是:", p)
}

type Person struct {
	Name string
}

func modifyStruct(p *Person) {
	fmt.Printf("函数内部的struct的内存地址是:%p\n", &p)
	p.Name = "张三"
}
