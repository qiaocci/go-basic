package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protobuf01/std"
)

func main() {
	res1 := &std.Result{
		Url:      "www.baidu.com",
		Title:    "百度上市了",
		Snippets: []string{"百度", "上市"},
	}
	res2 := &std.Result{
		Url:      "www.google.com",
		Title:    "谷歌上市了",
		Snippets: []string{"谷歌", "上市"},
	}

	var result []*std.Result
	result = append(result, res1, res2)

	student := &std.SearchResponse{
		Result: result,
	}
	data, err := proto.Marshal(student)
	if err != nil {
		log.Fatalf("Marshal failed, err:%v\n", err)
	}

	newStudent := &std.SearchResponse{}
	err = proto.Unmarshal(data, newStudent)
	if err != nil {
		log.Fatalf("UnMarshal failed, err:%v\n", err)
	}
	for idx, result := range newStudent.Result {
		fmt.Println(idx, result)
	}
}
