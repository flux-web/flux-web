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

const (
	workloadType    = "workload"
	helmreleaseType = "helmrelease"
	NotInRepo       = "NotInRepo"
)

var l = logs.GetLogger()

type Workload struct {
	ID         string `json:"ID"`
	Type       string
	Automated  bool
	Policies   interface{}
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

func NewWorkloads(data []byte, services Dictionary) []Workload {
	var w []Workload
	err := json.Unmarshal(data, &w)
	if err != nil {
		l.Panic(err.Error())
	}
	return initWorkloads(w, filterOutWorkloads(services))
}

func initWorkloads(workloads []Workload, services Dictionary) []Workload {
	for i, workload := range workloads {

		if services[workload.ID].Antecedent == "" {
			workloads[i].Type = workloadType
		} else {
			workloads[i].Type = helmreleaseType
		}

		workloads[i].Policies = services[workload.ID].Policies
		workloads[i].Automated = services[workload.ID].Automated

		for j, container := range workload.Containers {
			workloadStatus := MemGet(workload.getWorkloadKey(container.Name))
			if workloadStatus != "" {
				workloads[i].Containers[j].Status = workloadStatus
			} else {
				if container.Current.ID == container.Available[0].ID {
					workloads[i].Containers[j].Status = UpToDate
				} else {
					workloads[i].Containers[j].Status = Behind
				}
			}
		}
	}
	return workloads
}

func filterOutWorkloads(services Dictionary) Dictionary {
	for k, _ := range services {
		if services[k].Antecedent != "" || services[k].ReadOnly == NotInRepo {
			delete(services, k)
		}
	}
	return services
}
