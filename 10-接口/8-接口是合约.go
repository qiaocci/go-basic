package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	fmt.Printf("c=%v, type=%T, ptr=%p\n", c, c, c)
	//fmt.Printf("origin=%d\n", *c)
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	fmt.Printf("原始c指针=%p\n", &c)
	c.Write([]byte("hello"))
	fmt.Println("c.Write", c)

	var name = "dolly"
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println("fmt.Fprintf", c)
}
