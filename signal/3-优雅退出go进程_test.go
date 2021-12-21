package main

import (
	"os"
	"syscall"
	"testing"
)

func TestSignal(t *testing.T) {
	t.Log("int=", signum(syscall.SIGTERM))
}

func TestHandlerSet(t *testing.T) {
	sigTerm := 15
	tmp := sigTerm & 31
	t.Logf("tmp=%v, %T", tmp, tmp)
}

func BenchmarkOushu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OuShu2(1)
	}
}

// BenchmarkOushu-8   	1000000000	         0.3157 ns/op
func OuShu1(n int) bool {
	return n%2 == 0
}

// BenchmarkOushu-8   	1000000000	         0.3130 ns/op
func OuShu2(n int) bool {
	return n&1 == 0
}

func signum(sig os.Signal) int {
	switch sig := sig.(type) {
	case syscall.Signal:
		i := int(sig)
		if i < 0 || i >= 65 {
			return -1
		}
		return i
	default:
		return -1
	}
}
