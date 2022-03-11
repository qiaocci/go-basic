package main

import (
	"fmt"
	"image"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//print(Icon("left.png"))
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go Icon("left.png")
	}
	wg.Wait()
}

var icons map[string]image.Image

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	wg.Done()
	fmt.Println(icons[name])
	return icons[name]
}

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

func loadIcon(name string) image.Image {
	//time.Sleep(time.Second)
	return nil
}
