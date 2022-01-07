package main

import "fmt"

func main() {
	c := make(chan int)
	go read(c)

	c <- 1

	fmt.Println(c)
	close(c)
}
func write(c chan int) {
	c <- 1
}
func read(c chan int) {
	i := <-c // 阻塞
	fmt.Println(i)
}
