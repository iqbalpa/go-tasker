package main

import (
	"context"
	"fmt"
	"go-tasker/queue"
	"go-tasker/utils"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	
	// Define context deadline
	d := time.Now().Add(5 * time.Second)
	ctx, cancelCtx := context.WithDeadline(context.Background(), d)

	// Initialize 3 workers
	jq := queue.New(3)
	jq.StartWorkers(ctx, &wg)

	// Read the jobs
	jsonByte, err := utils.OpenFile("data/job.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the byte file to job object
	jobs, err := utils.Byte2Object(jsonByte)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add job to the queue (channel)
	for _, j := range jobs {
		jq.AddJob(
			ctx,
			j,
		)
	}

	// close the channel and wait for goroutines finished 
	close(jq.QueueCh) 
	wg.Wait()

	// cancel the context (optional), because I already defined the deadline
	cancelCtx()


	// ======================================================================
	fmt.Println(utils.CountCompleted(jobs))

	// convert object to byte
	jobsByte, err := utils.Object2Byte(jobs)
	if err != nil {
		fmt.Println(err)
		return
	}

	// save the jobs to a new file
	fname := "data/job-updated.json"
	err = utils.SaveFile(fname, jobsByte)
	if err != nil {
		fmt.Println(err)
		return
	}
}
