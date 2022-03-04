package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	Launch()
}

func Launch() {
	fmt.Println("lift off")
}
