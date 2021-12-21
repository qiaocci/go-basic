package main

import "fmt"

func main() {
	c := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-c:
			fmt.Println(x)
		case c <- i:
		}

	}
}
