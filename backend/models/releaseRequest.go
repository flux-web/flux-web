package models

import (
	"encoding/json"
)

type ReleaseRequest struct {
	Workload string `json:"Workload"`
	Container string `json:"Container"`
	Current   string `json:"Current"`
	Target    string `json:"Target"`
}

func NewReleseRequest(data []byte) (ReleaseRequest, error) {
	var r ReleaseRequest
	err := json.Unmarshal(data, &r)
	return r, err
}
