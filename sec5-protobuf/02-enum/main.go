package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protobuf01/std"
)

func main() {
	student := &std.Student{
		Name:   "qiaocc",
		Gender:   std.Student_MALE,
		Scores: []int32{98, 97, 96},
	}
	data, err := proto.Marshal(student)
	if err != nil {
		log.Fatalf("Marshal failed, err:%v\n", err)
	}

	newStudent := &std.Student{}
	err = proto.Unmarshal(data, newStudent)
	if err != nil {
		log.Fatalf("UnMarshal failed, err:%v\n", err)
	}
	fmt.Println(newStudent.Name, newStudent.Gender, newStudent.Scores)
}
