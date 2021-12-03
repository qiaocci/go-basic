package main

import (
	"syscall"
	"testing"
)

func Test(t *testing.T) {
	//var SIGINT    = Signal(0x2)
	// 16进制的2, 转成十进制, 也是2
	res := syscall.Signal(0x2)
	t.Log(res)
	t.Log(int(0x2))

}
