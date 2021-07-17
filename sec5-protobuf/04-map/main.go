package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protobuf01/std"
)

func main() {
	student := &std.MapRequest{
		Points: map[string]int32{
			"tom":   89,
			"linda": 95,
			"jerry": 78,
		},
	}
	data, err := proto.Marshal(student)
	if err != nil {
		log.Fatalf("Marshal failed, err:%v\n", err)
	}

	newStudent := &std.MapRequest{}
	err = proto.Unmarshal(data, newStudent)
	if err != nil {
		log.Fatalf("UnMarshal failed, err:%v\n", err)
	}

	fmt.Printf("req=%#v\n\n", newStudent)
	for k, v := range newStudent.Points {
		fmt.Println(k, v)
	}
}
