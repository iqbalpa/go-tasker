package job

import (
	"context"
	"fmt"
	"go-tasker/logger"
	"sync"
	"time"
)

func Process(ctx context.Context, wg *sync.WaitGroup, i int, ch chan *Job) {
	defer wg.Done()

	for {
		select {

		case <-ctx.Done():
			logger.Warn(
				fmt.Sprintf("Worker %d is exited", i),
			)
			return
			
		case job,ok := <-ch:
			if !ok {
				return
			}
			if job == nil {
				continue
			}

			logger.Info(
				fmt.Sprintf("Worker %d is processing job %d", i, job.Id),
			)
			
			job.Status = Running
			time.Sleep(1000 * time.Millisecond)
			job.Status = Completed
			
			logger.Info(
				fmt.Sprintf("Worker %d finished processing job %d", i, job.Id),
			)
		}
	}
}
