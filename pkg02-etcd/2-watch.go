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

	// watch
	ch := cli.Watch(context.Background(), "name") // <-chan WatchResponse
	for resp := range ch {
		for _, event := range resp.Events {
			fmt.Printf("type=%v, key=%s, value=%s\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}
