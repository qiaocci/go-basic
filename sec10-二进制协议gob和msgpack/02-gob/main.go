package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type s struct {
	data map[string]interface{}
}

func main() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1
	// encode
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(s1.data)
	if err != nil {
		log.Printf("gob encode failed, err:%v\n", err)
		return
	}
	b := buf.Bytes()
	fmt.Println(b)

	// decode
	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&s2.data)
	if err != nil {
		log.Printf("gob decode err:%v\n", err)
		return
	}
	fmt.Println(s2.data)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}

}
