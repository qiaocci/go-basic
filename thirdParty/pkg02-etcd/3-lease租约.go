package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	// 创建一个5s的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatalln(err)
	}
	// 5s后, tmpkey这个key就会被删除
	_, err = cli.Put(context.TODO(), "tmpkey", "tmpvalue", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatalln(err)
	}

	// 查看结果
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		resp, err := cli.Get(ctx, "tmpkey")
		if err != nil {
			return
		}
		cancel()
		if err != nil {
			fmt.Printf("put failed, err=%+v", err)
		}

		for _, item := range resp.Kvs {
			fmt.Printf("key=%s, value=%s\n", item.Key, item.Value)
		}

		time.Sleep(time.Second)
	}

}
