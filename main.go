package main

import (
	"go-tasker/queue"
)

func main() {
	jq := queue.New(3)
	jq.StartWorkers()
}
