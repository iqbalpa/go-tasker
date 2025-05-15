package main

import (
	"context"
	"fmt"
	"go-tasker/job"
	"go-tasker/queue"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	d := time.Now().Add(5 * time.Second)
	ctx, cancelCtx := context.WithDeadline(context.Background(), d)

	jq := queue.New(3)
	jq.StartWorkers(ctx, &wg)

	for i:=0; i<50; i++ {
		jq.AddJob(
			string(job.Email),
			fmt.Sprintf("Message %d", i+1),
			i+1,
		)
	}

	close(jq.QueueCh)
	wg.Wait()

	cancelCtx()
}
