package job

type Job struct {
	Id      int       `json:'id'`
	Payload string    `json:'payload'`
	JobType JobType   `json:'jobType'`
	Status  JobStatus `json:'status'`
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
