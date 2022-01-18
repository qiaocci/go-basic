package main

import (
	"log"
)

func main() {
	counters := make(chan int)
	pingFangs := make(chan int)

	// pipeline 1, 添加数据
	go counter(counters)

	// pipeline 2, 平方
	go pingfang(counters, pingFangs)

	// pipeline 3, 平方
	printNum(pingFangs)
}

func counter(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
	close(out)
}

func pingfang(out <-chan int, in chan<- int) {
	for item := range out {
		in <- item * item
	}
	close(in)
}

func printNum(in <-chan int) {
	for item := range in {
		log.Println(item)
	}
}
