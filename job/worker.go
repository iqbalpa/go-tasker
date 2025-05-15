package job

import (
	"fmt"
	"go-tasker/logger"
	"time"
)

func Process(i int, ch chan *Job) {
	// this is infinite loop
	// it will always looping over jobs in the channel
	for job := range ch {
		logger.Info(
			fmt.Sprintf("Worker %d is processing job %d", i, job.Id),
		)
		
		job.Status = Running
		timer := time.NewTimer(100 * time.Millisecond)
		<-timer.C
		job.Status = Completed

		logger.Info(
			fmt.Sprintf("Worker %d finished processing job %d", i, job.Id),
		)
	}
}
