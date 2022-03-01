package main

import (
	"go-basic/12-concurrency/22-bingfa-xunhuan/thumbnail"
	"log"
)

func main() {
	images := []string{
		"034dda6837dffc71caf0f7e43b996f09.jpeg",
		"1c67f2e21fca0b49dcc06ae25dd88d1c.jpeg",
		"29ce370698955bfe0a5e831653478470.jpeg",
		"579e31170c04c472a584958465b12b41.jpeg",
		"c03782e80c8beaa4a4d58e5dab4222da.jpeg",
		"dcf25321367bc43cc7851ed300efd79a.jpeg",
		"f268cff89a2cd9f025ef5d73122b1b96.jpeg",
		"fa8efbd17e9c741497e512b231f8ba68.jpeg",
		"photo_2021-08-19_14-18-45.jpg",
	}
	var filenames []string
	for _, image := range images {
		filenames = append(filenames, "/Users/qiaocc/Pictures/image/"+image)
	}

	makeThumbnails4(filenames)
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

// 等待goroutine完成
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	// wait goroutine to complete
	for range filenames {
		<-ch
	}
}

// 处理错误
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	// wait goroutine to complete
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}
