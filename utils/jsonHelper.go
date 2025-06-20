package utils

import (
	"encoding/json"
	"fmt"
	"go-tasker/job"
	"io"
	"os"
)

func OpenFile(fname string) ([]byte, error) {
	// open the file
	jsonFile, err := os.Open(fname)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to open json file")
	}
	defer jsonFile.Close()
	
	// read the byte
	jsonByte, err := io.ReadAll(jsonFile)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to open json file")
	}
	
	return jsonByte, nil
}

func Byte2Object(jsonByte []byte) ([]*job.Job, error) {
	var jobs []*job.Job
	
	err := json.Unmarshal(jsonByte, &jobs)
	if err != nil {
		return []*job.Job{}, fmt.Errorf("failed to decode the json")
	}

	return jobs, nil
}

func CountCompleted(jobs []*job.Job) (int, int) {
	completed := 0
	pending := 0

	for _, j := range jobs {
		if j.Status == job.Completed {
			completed += 1
		} else {
			pending += 1
		}
	}

	return completed, pending
}

func Object2Byte(jobs []*job.Job) ([]byte, error) {
	b, err := json.Marshal(jobs)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to encode the jobs")
	}
	return b, nil
}

func SaveFile(fname string, byteJobs []byte) (error) {
	err := os.WriteFile(fname, byteJobs, 0644)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to save the file")
	}
	return nil
}
