package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() { // 开启goroutine， 将0-100放入c1
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() { // 开启goroutine， 从c1中取数据，放入c2
		for { // 从通道取值方法1
			i, ok := <-c1
			if !ok {
				break
			}
			c2 <- i * i
		}
		close(c2)
	}()

	for i := range c2 { // 从通道取值，方法2
		fmt.Println(i)
	}
}
