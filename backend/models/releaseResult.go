package models

import (
	"encoding/json"
)

type ReleaseResult struct {
	Status string 
	Workload string
	Container string
	Tag string
}

func NewReleseResult(data []byte) (ReleaseResult, error) {
	var r ReleaseResult
	err := json.Unmarshal(data, &r)
	return r, err
}
