package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(time.Second * 5):
	case <-abort:
		fmt.Println("aborted")
		return
	}
	launch()
}

func launch() {
	fmt.Println("lift off")
}
