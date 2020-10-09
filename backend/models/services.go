package models

import (
	"encoding/json"
	"log"
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
		automated     string `json:"automated"`
		tagChartImage string `json:"tag.chart-image"`
	} `json:"Policies"`
}

func NewServices(data []byte) ([]Service, error) {
	var s []Service
	log.Printf("JSON: %v", string(data))
	err := json.Unmarshal(data, &s)
	for _, value := range s {
		log.Printf("Policies: %v / %v", value.Policies.automated, value.Policies.tagChartImage)
	}
	return s, err
}
