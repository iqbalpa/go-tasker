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

func Byte2Object(jsonByte []byte) ([]job.Job, error) {
	var jobs []job.Job
	
	err := json.Unmarshal(jsonByte, &jobs)
	if err != nil {
		return []job.Job{}, fmt.Errorf("failed to decode the json")
	}

	return jobs, nil
}