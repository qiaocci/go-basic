package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()

	// set
	conn.Do("set", "name", "qiaocc")
	// get
	reply, _ := redis.String(conn.Do("get", "name"))
	log.Println(reply)

}
