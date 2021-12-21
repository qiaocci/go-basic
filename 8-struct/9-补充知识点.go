package main

import "fmt"

type Person1 struct {
	name   string
	age    int
	dreams []string
}

func (p *Person1) setDreams(dreams []string) {
	fmt.Printf("in function: %p\n", &p)
	p.dreams = dreams
	//p.dreams = make([]string, len(dreams))
	//copy(p.dreams, dreams)
}

func main() {
	p1 := Person1{
		name: "qiaocc",
		age:  21,
	}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	fmt.Printf("before function: %p\n", &p1)

	p1.setDreams(data)
	fmt.Println(p1, data)

	data[0] = "不吃饭" // 会修改data和p中的值
	fmt.Println(p1, data)

}
