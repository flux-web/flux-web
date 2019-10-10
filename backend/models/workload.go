package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/logs"
)

const (
	UpToDate string = string(iota + 'a')
	Behind
	InRelease
	ReleaseFaild
)

var l = logs.GetLogger()

type Workload struct {
	ID         string `json:"ID"`
	Status     string
	Containers []struct {
		Name    string `json:"Name"`
		Current struct {
			ID      string `json:"ID"`
			Digest  string `json:"Digest"`
			ImageID string `json:"ImageID"`
			Labels  struct {
				OrgLabelSchemaBuildDate       time.Time `json:"org.label-schema.build-date"`
				OrgOpencontainersImageCreated time.Time `json:"org.opencontainers.image.created"`
			} `json:"Labels"`
			CreatedAt   time.Time `json:"CreatedAt"`
			LastFetched time.Time `json:"LastFetched"`
		} `json:"Current"`
		Available []struct {
			ID      string `json:"ID"`
			Digest  string `json:"Digest"`
			ImageID string `json:"ImageID"`
			Labels  struct {
				OrgLabelSchemaBuildDate       time.Time `json:"org.label-schema.build-date"`
				OrgOpencontainersImageCreated time.Time `json:"org.opencontainers.image.created"`
			} `json:"Labels,omitempty"`
			CreatedAt   time.Time `json:"CreatedAt"`
			LastFetched time.Time `json:"LastFetched"`
		} `json:"Available"`
		AvailableImagesCount int `json:"AvailableImagesCount"`
		FilteredImagesCount  int `json:"FilteredImagesCount"`
	} `json:"Containers"`
}

func (w Workload) getWorkloadKey() string {
	return w.ID + "_" + w.Containers[0].Name
}

func NewWorkloads(data []byte) []Workload {
	var w []Workload
	err := json.Unmarshal(data, &w)
	if err != nil {
		l.Panic(err.Error)
	}
	return setWorkloadsStatus(w)
}

func setWorkloadsStatus(workloads []Workload) []Workload {
	for _, workload := range workloads {
		workloadStatus := MemGet(workload.getWorkloadKey())
		if workloadStatus != "" {
			workload.Status = workloadStatus
		} else {
			for _, container := range workload.Containers {
				if container.Current.ID == container.Available[0].ID {
					workload.Status = UpToDate
				} else {
					workload.Status = Behind
				}
			}
		}
		l.Println(workload.ID + ": " + workload.Status)
	}
	return workloads
}
