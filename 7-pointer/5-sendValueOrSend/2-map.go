package main

import "fmt"

// func makemap(t *maptype, hint int, h *hmap) *hmap
// func makechan(t *chantype, size int) *hchan
func main() {
	persons := make(map[string]int)
	persons["张三"] = 18

	mp := &persons
	fmt.Printf("原始map的内存地址是:%p\n", mp)
	modifyMap(persons)
	fmt.Println("map的值被修改了, 新值是:", persons)
}

func modifyMap(p map[string]int) {
	fmt.Printf("函数内部map的内存地址是:%p\n", p)
	p["张三"] = 20
}
