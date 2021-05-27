package main

import (
	"fmt"
	"html"
	"net/http"
)

func fooHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello qiaocc")
}

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello %q", html.EscapeString(request.URL.Path))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}
