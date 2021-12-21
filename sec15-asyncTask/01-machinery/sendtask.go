package main

import (
	"15-async/01-machinery/worker"
	"context"
)

func main() {
	worker.SendHelloWorldTask(context.Background())
}
