package main

import "sync"

type singleton struct {
}

var instance *singleton
var once sync.Once

// 并发安全的单例模式
func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
