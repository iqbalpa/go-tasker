package queue

import (
	"fmt"
	"go-tasker/job"
	"go-tasker/logger"
	"sync"
)

type JobQueue struct {
	mu sync.Mutex
	PendingJob []*job.Job
	AllJob []*job.Job
	NumWorker int
	QueueCh chan *job.Job
}

// channel is where the worker look up to 
func New(numWorker int) *JobQueue {
	return &JobQueue{
		PendingJob: []*job.Job{},
		AllJob: []*job.Job{},
		NumWorker: numWorker,
		mu : sync.Mutex{},
		QueueCh: make(chan *job.Job),
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
	jq.PendingJob = append(jq.PendingJob, &j)
	return "Added new job successfully", nil
}

func (jq *JobQueue) RetrieveJob() (*job.Job, error) {
	if len(jq.PendingJob) == 0 {
		return nil, fmt.Errorf("there is no job right now")
	}
	jq.mu.Lock()
	defer jq.mu.Unlock()

	job := jq.PendingJob[0]
	jq.PendingJob = jq.PendingJob[1:]
	return job, nil
}

func (jq *JobQueue) StartWorkers() {
	logger.Info("Starting the workers!")
	// create workers (goroutines)
	for i := 0; i < jq.NumWorker; i++ {
		logger.Info(fmt.Sprintf("Creating worker %d", i+1))
		go job.Process(i, jq.QueueCh)
	}
	// infinite loop to retrieve job 
	// and pass it into channel
	for {
		j, err := jq.RetrieveJob()
		if err != nil {
			continue
		}
		jq.QueueCh <- j
	}
}
