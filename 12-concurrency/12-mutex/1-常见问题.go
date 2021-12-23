package main

import (
	"fmt"
	"sync"
)

type T struct {
	lock sync.Mutex
}

func (t *T) Lock() {
	t.lock.Lock()
}

// 报错(t T) Unlock()
// 这种也就是 Unlock(t), 因为是值传递, t.lock被复制了一份. 这样就会报错了
// 使用go vet 1-常见问题.go, 报错: Unlock passes lock by value
// 相关链接: https://medium.com/golangspec/detect-locks-passed-by-value-in-go-efb4ac9a3f2b
func (t *T) Unlock() {
	fmt.Println("函数内")
	fmt.Printf("t的指针=%p\n", t)            // 0xc0000b4008
	fmt.Printf("t.lock的指针=%p\n", &t.lock) // 0xc0000b4008
	t.lock.Unlock()
}
func main() {
	t := T{lock: sync.Mutex{}}
	fmt.Println("函数外")
	fmt.Printf("t的指针=%p\n", &t)           // 0xc0000b4008
	fmt.Printf("t.lock的指针=%p\n", &t.lock) // 0xc0000b4008

	t.Lock() // 底层是: (&t).Lock()
	t.Unlock()
	t.Lock()
}
