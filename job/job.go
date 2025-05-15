package job

import "sync"

type Job struct {
	sync.Mutex
	Id      int
	Payload string
	JobType JobType
	Status  JobStatus
}

type JobType string

const (
	Email    JobType = "email"
	Reminder JobType = "remainder"
)

type JobStatus string

const (
	Pending   JobStatus = "pending"
	Running   JobStatus = "running"
	Completed JobStatus = "completed"
)
