package models

import (
	"encoding/json"
)

type ServiceDictionary map[string]Service

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

func NewServices(data []byte) (ServiceDictionary, error) {
	var s []Service
	serviceDictionary := ServiceDictionary{}
	err := json.Unmarshal(data, &s)

	for i, service := range s {
		serviceDictionary[service.ID] = s[i]
	}
	return serviceDictionary, err
}
