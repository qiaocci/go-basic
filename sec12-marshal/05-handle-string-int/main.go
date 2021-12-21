package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Card struct {
	ID    int64   `json:"id,string"` // 添加string tag
	Score float64 `json:"score"`
}

func main() {
	intAndString()
}

func intAndString() {
	jsonStr := `{"id": "123", "score": 89.11}`
	var c Card
	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		log.Printf("json unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c: %#v, id=%v, type id=%T\n", c, c.ID, c.ID)

	//num, err := strconv.Atoi(c.ID)
	//if err != nil {
	//	println(err)
	//	return
	//}
	//
	//fmt.Printf("num=%v, type=%T\n", num, num)
}
