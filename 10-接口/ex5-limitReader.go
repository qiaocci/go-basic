package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "hi there"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 4)
	n, err := b.ReadFrom(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(b.String())
	fmt.Println(b.String() == "hi t")
}

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}
