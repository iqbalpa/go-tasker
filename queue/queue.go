package queue

import (
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
		mu : sync.Mutex{},
		QueueCh: make(chan *job.Job, 50),
	}
}

func (jq *JobQueue) ListJob() (*[]*job.Job) {
	return &jq.AllJob
}

func (jq *JobQueue) AddJob(typeStr string, payload string, id int) (string, error) {
	j := job.Job{
		Payload: payload,
		JobType: job.JobType(typeStr),
		Status: job.Pending,
		Id: id,
	}
	jq.mu.Lock()
	defer jq.mu.Unlock()

	jq.AllJob = append(jq.AllJob, &j)
	jq.QueueCh <- &j
	return "Added new job successfully", nil
}

func (jq *JobQueue) StartWorkers() {
	logger.Info("Starting the workers!")
	// create workers (goroutines)
	for i := 0; i < jq.NumWorker; i++ {
		logger.Info(fmt.Sprintf("Creating worker %d", i+1))
		go job.Process(i, jq.QueueCh)
	}
}
