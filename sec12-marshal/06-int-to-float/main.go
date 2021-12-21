package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	jsonDemo()
}
func jsonDemo() {
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str: %#v\n", string(b))

	var m2 map[string]interface{}
	err = json.Unmarshal(b, &m2)
	if err != nil {
		fmt.Printf("unmrshal failed, err:%v\n", err)
	}
	fmt.Printf("m2=%#v, count=%v, count type=%T\n", m2, m2["count"], m2["count"])

	s := "2021-06-28 16:17:00"
	formatDateWithoutSlash := "20060102"

	//t, _ :=time.Parse(formatDateWithoutSlash, s)
	//m3 :=map[string]interface{}{
	//	"i": s,
	//}
	layout := "2006-01-02 16:04:05"

	t, err := time.Parse(s, layout)
	fmt.Printf("%v%03d", t.Format(formatDateWithoutSlash), 2)

	//res :=fmt.Sprintf("%v%03d", )
	//fmt.Printf("res=%v\n", res)

}
