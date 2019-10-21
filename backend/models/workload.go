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
	return initWorkloads(w, services)
}

func initWorkloads(workloads []Workload, services ServiceDictionary) []Workload {

	filterdWorkloads := []Workload{}

	//remove workloads that belongs to a chart
	removeChildWorkloads(&services)

	for i, workload := range workloads {

		if services[workload.ID].ID == "" {
			remove(workloads, i)
			continue
		}

		setType(workload.ID, &workloads[i])
		workloads[i].Automated = services[workload.ID].Automated
		initContainers(&workloads[i], services)
		filterdWorkloads = append(filterdWorkloads, workloads[i])
	}
	return filterdWorkloads
}

func initContainers(workload *Workload, services ServiceDictionary) {
	for i, container := range workload.Containers {
		if services[workload.ID].Policies != nil {
			workload.Policies = services[workload.ID].Policies
			workload.Containers[i].Available = filterOutTags(services[workload.ID].Policies, workload.Containers[i].Available, workload.Containers[i].Name)
		}

		workloadStatus := MemGet(workload.getWorkloadKey(container.Name))

		if workloadStatus != "" {
			workload.Containers[i].Status = workloadStatus
		} else {
			if len(container.Available) > 0 && container.Current.ID == container.Available[0].ID {
				workload.Containers[i].Status = UpToDate
			} else {
				workload.Containers[i].Status = Behind
			}
		}
	}
}

func setType(workloadID string, workload *Workload) {
	if strings.Contains(workloadID, helmreleaseType) {
		workload.Type = helmreleaseType
	} else {
		workload.Type = workloadType
	}
}

func removeChildWorkloads(services *ServiceDictionary) {
	for k, service := range *services {
		if service.Antecedent != "" {
			delete(*services, k)
		}
	}
}

func filterOutTags(policies interface{}, available []Available, containerName string) []Available {
	policiesMap := policies.(map[string]interface{})
	//filterdAvailable := []Available{}
	l.Println(containerName)
	for k, v := range policiesMap {
		availableName := strings.Split(k, ".")

		// not a tag.<name>
		if len(availableName) == 1 {
			continue
		}
		tagPattern := strings.Split(fmt.Sprintf("%s", v), ":")
		if tagPattern[0] != "glob" || len(tagPattern[1]) == 1 {
			continue
		}

		l.Println("tagFilter: " + tagPattern[1])
		if containerName == availableName[1] {
			i := 0
			for _, ava := range available {
				matched, err := regexp.MatchString(tagPattern[1], strings.Split(ava.ID, ":")[1])

				if err != nil {
					panic(err.Error())
				}
				if matched {
					available[i] = ava
					i++
				}
			}
			available = available[:i]
		}
	}
	return available
}

func remove(s []Workload, i int) []Workload {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
