package models

import (
	"encoding/json"
)

type Service struct {
	ID         string `json:"ID"`
	Containers []struct {
		Name    string `json:"Name"`
		Current struct {
			ID     string `json:"ID"`
			Labels struct {
			} `json:"Labels"`
		} `json:"Current"`
		LatestFiltered struct {
			ID     string `json:"ID"`
			Labels struct {
			} `json:"Labels"`
		} `json:"LatestFiltered"`
	} `json:"Containers"`
	ReadOnly  string `json:"ReadOnly"`
	Status    string `json:"Status"`
	Automated bool   `json:"Automated"`
	Policies  struct {
		Automated     string `json:"automated"`
		TagChartImage string `json:"tag.chart-image"`
	} `json:"Policies"`
}

func NewServices(data []byte) ([]Service, error) {
	var s []Service
	err := json.Unmarshal(data, &s)
	return s, err
}
