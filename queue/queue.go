package queue

import (
	"context"
	"fmt"
	"go-tasker/job"
	"go-tasker/logger"
	"sync"
)

type JobQueue struct {
	mu sync.Mutex
	AllJob []*job.Job
	NumWorker int
	QueueCh chan *job.Job
}

// channel is where the worker look up to 
func New(numWorker int) *JobQueue {
	return &JobQueue{
		AllJob: []*job.Job{},
		NumWorker: numWorker,
		QueueCh: make(chan *job.Job, 50),
		mu : sync.Mutex{},
	}
}

func (jq *JobQueue) ListJob() ([]*job.Job) {
	return jq.AllJob
}

func (jq *JobQueue) AddJob(ctx context.Context, j *job.Job) (string, error) {
	select {
	case <-ctx.Done():
		logger.Error(fmt.Sprintf("failed to queue job %d: context deadline exceeded", j.Id))
		return "", ctx.Err()
	case jq.QueueCh <-j:
		jq.mu.Lock()
		defer jq.mu.Unlock()
		jq.AllJob = append(jq.AllJob, j)
	}

	return "Added new job successfully", nil
}

func (jq *JobQueue) StartWorkers(ctx context.Context, wg *sync.WaitGroup) {
	logger.Info("Starting the workers!")

	// create workers (goroutines)
	for i := 0; i < jq.NumWorker; i++ {
		logger.Info(fmt.Sprintf("Creating worker %d", i+1))
		wg.Add(1)
		go job.Process(ctx, wg, i, jq.QueueCh)
	}
}
