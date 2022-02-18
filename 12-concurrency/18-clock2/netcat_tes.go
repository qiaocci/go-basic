package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	src := strings.NewReader("hello\n")
	dst := os.Stdout
	bytes, err := io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the number of bytes are:%v\n", bytes)
}
