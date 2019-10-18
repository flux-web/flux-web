package models

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
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
	Containers []Container
}

type Container struct {
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
	Available            []Available
	AvailableImagesCount int `json:"AvailableImagesCount"`
	FilteredImagesCount  int `json:"FilteredImagesCount"`
}

type Available struct {
	ID      string `json:"ID"`
	Digest  string `json:"Digest"`
	ImageID string `json:"ImageID"`
	Labels  struct {
		OrgLabelSchemaBuildDate       time.Time `json:"org.label-schema.build-date"`
		OrgOpencontainersImageCreated time.Time `json:"org.opencontainers.image.created"`
	} `json:"Labels,omitempty"`
	CreatedAt   time.Time `json:"CreatedAt"`
	LastFetched time.Time `json:"LastFetched"`
}

func (w Workload) getWorkloadKey(container string) string {
	return w.ID + "_" + container
}

func NewWorkloads(data []byte, services ServiceDictionary) []Workload {
	var w []Workload
	err := json.Unmarshal(data, &w)
	if err != nil {
		l.Panic(err.Error())
	}
	for k, _ := range services {
		l.Println("@@@@@@@")
		l.Println(services[k].ID)
		l.Println(services[k].Antecedent)
	}
	return initWorkloads(w, services)
}

func initWorkloads(workloads []Workload, services ServiceDictionary) []Workload {

	filterdWorkloads := []Workload{}

	//remove workloads that belongs to a chart
	for k, _ := range services {
		if services[k].Antecedent != "" {
			delete(services, k)
		}
	}

	for i, workload := range workloads {

		if services[workload.ID].ID == "" {
			remove(workloads, i)
			continue
		}

		if strings.Contains(workload.ID, helmreleaseType) {
			workloads[i].Type = helmreleaseType
		} else {
			workloads[i].Type = workloadType
		}

		workloads[i].Automated = services[workload.ID].Automated

		for j, container := range workload.Containers {
			if services[workload.ID].Policies != nil {
				workloads[i].Policies = services[workload.ID].Policies
				workloads[i].Containers[j].Available = filterOutTags(services[workload.ID].Policies, workloads[i].Containers[j].Available, workloads[i].Containers[j].Name)
			}

			workloadStatus := MemGet(workload.getWorkloadKey(container.Name))
			if workloadStatus != "" {
				workloads[i].Containers[j].Status = workloadStatus
			} else {
				if len(container.Available) > 0 && container.Current.ID == container.Available[0].ID {
					workloads[i].Containers[j].Status = UpToDate
				} else {
					workloads[i].Containers[j].Status = Behind
				}
			}
		}
		filterdWorkloads = append(filterdWorkloads, workloads[i])
	}
	return filterdWorkloads
}

func filterOutTags(policies interface{}, available []Available, containerName string) []Available {
	policiesMap := policies.(map[string]interface{})
	l.Println(containerName)
	for k, v := range policiesMap {
		availableName := strings.Split(k, ".")
		if len(availableName) == 1 {
			continue
		}
		tagPattern := strings.Split(fmt.Sprintf("%s", v), ":")
		if tagPattern[0] != "glob" {
			continue
		}
		l.Println("tagFilter: " + tagPattern[1])
		if containerName == availableName[1] {
			for _, ava := range available {
				re := regexp.MustCompile(`\*`)

				if !re.MatchString(strings.Split(ava.ID, ":")[1]) {
					l.Println("need to delete " + strings.Split(ava.ID, ":")[1])
				}
			}
		}
	}
	return available
}

func remove(s []Workload, i int) []Workload {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
