package main

import (
	"bytes"
	"strings"
	"testing"
)

// 字符串拼接
// https://dev.to/shiraazm/efficiently-concatenating-strings-in-go-6do
func BenchmarkStringConcatenation(b *testing.B) {
	benchmarks := []struct {
		name   string
		testFn func(int)
		param  int
	}{
		{"Regular 50", regularConcatentation, 50},
		{"Regular 100", regularConcatentation, 100},
		{"bytes.Buffer 50", bytesBuffer, 100},
		{"bytes.Buffer 100", bytesBuffer, 100},
		{"strings.Builder 50", stringBuilder, 100},
		{"strings.Builder 100", stringBuilder, 100},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.testFn(bm.param)
			}
		})
	}
}

func BenchmarkRegularConcatenation40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regularConcatentation(40)
	}
}

func BenchmarkBytesBuffer40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesBuffer(40)
	}
}

func BenchmarkStringBuilder40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilder(40)
	}
}

func bytesBuffer(concats int) {
	var s bytes.Buffer
	for i := 0; i < concats; i++ {
		s.WriteString("somestring")
	}
}

func regularConcatentation(concats int) {
	s := ""
	for i := 0; i < concats; i++ {
		s += "somestring"
	}
}

func stringBuilder(concats int) {
	var s strings.Builder
	for i := 0; i < concats; i++ {
		s.WriteString("somestring")
	}
}
