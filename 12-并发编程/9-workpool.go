package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500) // 休息500毫秒
		fmt.Printf("worker:%d stop job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 200)

	// 开启3个goroutine
	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}
	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//输出结果
	//for ret := range results {
	//	fmt.Println(ret)
	//}

	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
