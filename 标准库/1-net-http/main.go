package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//get()
	//getWithParams()
	post()
}

func get() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

func getWithParams() {
	apiUrl := "http://httpbin.org/get"
	data := url.Values{}
	data.Set("name", "qiaocc")
	data.Set("age", "29")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url url failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

func post()  {
	apiUrl := "http://httpbin.org/post"
	contentType:="application/json"
	data:=`{"name": "qiaocc", "age": 29}`
	resp, err:=http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}