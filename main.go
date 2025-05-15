package main

import (
	"fmt"
	"go-tasker/job"
	"go-tasker/queue"
)

func main() {
	jq := queue.New(3)
	jq.StartWorkers()

	for i:=0; i<100; i++ {
		jq.AddJob(
			string(job.Email),
			fmt.Sprintf("Message %d", i+1),
			i+1,
		)
	}
}
