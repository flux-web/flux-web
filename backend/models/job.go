package models

import (
	"encoding/json"
)

type Job struct {
	Result struct {
		Revision string `json:"revision"`
		Spec     struct {
			Type  string `json:"type"`
			Cause struct {
				Message string `json:"Message"`
				User    string `json:"User"`
			} `json:"cause"`
			Spec struct {
				Kind           string `json:"Kind"`
				ContainerSpecs struct {
					Deployment []struct {
						Container string `json:"Container"`
						Current   string `json:"Current"`
						Target    string `json:"Target"`
					} `json:"deployment"`
				} `json:"ContainerSpecs"`
				SkipMismatches bool `json:"SkipMismatches"`
				Force          bool `json:"Force"`
			} `json:"spec"`
		} `json:"spec"`
		Result struct {
			Deployment struct {
				Status       string      `json:"Status"`
				Error        string      `json:"Error"`
				PerContainer interface{} `json:"PerContainer"`
			} `json:"deploymentt"`
		} `json:"result"`
	} `json:"Result"`
	Err          string `json:"Err"`
	StatusString string `json:"StatusString"`
}

func NewJob(data []byte) (Job, error) {
	var j Job
	err := json.Unmarshal(data, &j)
	return j, err
}
