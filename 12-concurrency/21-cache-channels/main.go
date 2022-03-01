package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ch := make(chan string, 3)
	go func() { ch <- request("https://www.google.com") }()
	go func() { ch <- request("https://www.youtube.com") }()
	go func() { ch <- request("https://www.baidu.com") }()

	fmt.Println(<-ch)
}

func request(hostname string) (response string) {
	resp, err := http.Get(hostname)
	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
