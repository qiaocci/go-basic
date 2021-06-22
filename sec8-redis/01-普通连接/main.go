package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var rdb *redis.Client

func main() {
	V8Example()
}

func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		log.Print(err)
		return
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, // 连接池大小
	})
	//最新版本的go-redis库的相关命令都需要传递context.Context参数
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("ping err:%v\n", err)
		return
	}
	return nil
}
