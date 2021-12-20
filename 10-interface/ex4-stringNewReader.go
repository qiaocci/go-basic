package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	s := "hi there"
	b := &bytes.Buffer{}
	b.ReadFrom(NewReader(s))
	fmt.Println(b.String() == s)
}

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}
