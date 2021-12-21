package main

import "15-async/01-machinery/worker"

func main() {
	// 启动异步任务框架
	taskWorker := worker.NewAsyncTaskWorker(0)
	taskWorker.Launch()
}
