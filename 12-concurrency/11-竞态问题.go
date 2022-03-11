package main

import (
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup
var x int

func add() {
	for i := 0; i < 5000; i++ {
		x = x + i
		//time.Sleep(time.Millisecond * 1)
	}
	wg3.Done()
}

func main() {
	wg3.Add(2)
	go add()
	go add()
	wg3.Wait()
	fmt.Println(x) // 正确值： 24995000， 但是打印的值有可能不对。
}
