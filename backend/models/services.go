package models

import (
	"encoding/json"
)

type Dictionary map[string]Service

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
	ReadOnly string `json:"ReadOnly"`
	Status   string `json:"Status"`
	Rollout  struct {
		Desired   int         `json:"Desired"`
		Updated   int         `json:"Updated"`
		Ready     int         `json:"Ready"`
		Available int         `json:"Available"`
		Outdated  int         `json:"Outdated"`
		Messages  interface{} `json:"Messages"`
	} `json:"Rollout"`
	SyncError  string      `json:"SyncError"`
	Antecedent string      `json:"Antecedent"`
	Labels     struct{}    `json:"Labels"`
	Automated  bool        `json:"Automated"`
	Locked     bool        `json:"Locked"`
	Ignore     bool        `json:"Ignore"`
	Policies   interface{} `json:"Policies"`
}

func NewServices(data []byte) (Dictionary, error) {
	var s []Service
	serviceDictionary := Dictionary{}
	err := json.Unmarshal(data, &s)

	for _, service := range s {
		serviceDictionary[service.ID] = service
	}

	return serviceDictionary, err
}
