package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
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

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "name", "tom")
	cancel()
	if err != nil {
		fmt.Printf("put failed, err=%+v", err)
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "name")
	cancel()
	if err != nil {
		fmt.Printf("put failed, err=%+v", err)
	}

	for _, item := range resp.Kvs {
		fmt.Printf("key=%s, value=%s\n", item.Key, item.Value)
	}
}
