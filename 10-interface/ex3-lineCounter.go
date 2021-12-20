package main

import (
	"fmt"
)

func main() {
	c := &LineCounter{}
	c.Writer([]byte("你好\n世界\n2021\n12月"))
	fmt.Println(c)
	fmt.Println(c.N())
}

type LineCounter struct {
	lines int
}

func (c *LineCounter) Writer(p []byte) (int, error) {
	for _, line := range p {
		if line == '\n' {
			c.lines++
		}
	}

	return c.lines, nil
}

func (c *LineCounter) String() string {
	return fmt.Sprintf("%d", c.lines)
}

func (c *LineCounter) N() int {
	return c.lines
}
