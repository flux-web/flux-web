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
	Containers []struct {
		Name    string `json:"Name"`
		Status  string
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

func (w Workload) getWorkloadKey(container string) string {
	return w.ID + "_" + container
}

func NewWorkloads(data []byte) []Workload {
	var w []Workload
	err := json.Unmarshal(data, &w)
	if err != nil {
		l.Panic(err.Error())
	}
	return setWorkloadsStatus(w)
}

func setWorkloadsStatus(workloads []Workload) []Workload {
	l.Printf("in setWorkloadsStatus")
	for i, workload := range workloads {
		for j, container := range workload.Containers {
			workloadStatus := MemGet(workload.getWorkloadKey(container.Name))
			if workloadStatus != "" {
				workloads[i].Containers[j].Status = workloadStatus
			} else {
				if container.Current.ID == container.Available[0].ID {
					workloads[i].Containers[j].Status = Behind
				} else {
					workloads[i].Containers[j].Status = UpToDate
				}
			}
		}
	}
	return workloads
}
