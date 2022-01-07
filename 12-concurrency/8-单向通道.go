package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go f1(ch1)
	go f2(ch1, ch2)

	for res := range ch2 {
		println(res)
	}
}

// 代表ch只能接收值
func f1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//<-ch // 报错, 不能发送值
	close(ch)
}

// ch1只能发送值, ch2只能接收值
func f2(ch1 <-chan int, ch2 chan<- int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}
