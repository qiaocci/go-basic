package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type respData struct {
	resp *http.Response
	err  error
}

func main() {
	// 定义100毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel() // 调用cancel释放子goutine的资源
	doCall(ctx)
}

func doCall(ctx context.Context) {
	// 请求频繁可定义全局的client对象并启用长链接
	// 请求不频繁使用短链接
	// https://www.jianshu.com/p/14edebc91c1b
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{Transport: &transport}

	respChan := make(chan *respData)

	req, err := http.NewRequest("GET", "http://127.0.0.1:8000", nil)
	if err != nil {
		fmt.Printf("ner request failed, err=%+v\n", err)
		return
	}
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client do, resp=%v, err=%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()
	defer wg.Wait()
	select {
	case <-ctx.Done():
		fmt.Println("请求超时")

	case result := <-respChan:
		if result.err != nil {
			fmt.Printf("call server api failed, err=%+v\n", result.err)
			return
		}
		all, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp=%v\n", string(all))
	}
}
