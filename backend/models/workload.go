package models

import (
	"log"
	"regexp"
	"strings"
	"time"
)

type Workload struct {
	ID         string      `json:"ID"`
	Containers []Container `json:"Containers"`
	ReadOnly   string      `json:"ReadOnly"`
	Status     string      `json:"Status"`
	Automated  bool        `json:"Automated"`
	Policies   struct {
		Automated     string `json:"automated"`
		TagChartImage string `json:"tag.chart-image"`
	} `json:"Policies"`
}

type Container struct {
	Name    string `json:"Name"`
	Current struct {
		ID          string    `json:"ID"`
		Digest      string    `json:"Digest"`
		ImageID     string    `json:"ImageID"`
		Labels      Labels    `json:"Labels"`
		CreatedAt   time.Time `json:"CreatedAt"`
		LastFetched time.Time `json:"LastFetched"`
	} `json:"Current"`
	LatestFiltered struct {
		ID          string    `json:"ID"`
		Digest      string    `json:"Digest"`
		ImageID     string    `json:"ImageID"`
		Labels      Labels    `json:"Labels"`
		CreatedAt   time.Time `json:"CreatedAt"`
		LastFetched time.Time `json:"LastFetched"`
	} `json:"LatestFiltered"`
	Available            []Available `json:"Available"`
	AvailableImagesCount int         `json:"AvailableImagesCount"`
	FilteredImagesCount  int         `json:"FilteredImagesCount"`
}

type Available struct {
	ID          string    `json:"ID"`
	Digest      string    `json:"Digest"`
	ImageID     string    `json:"ImageID"`
	Labels      Labels    `json:"Labels"`
	CreatedAt   time.Time `json:"CreatedAt"`
	LastFetched time.Time `json:"LastFetched"`
}

type Labels struct {
	OrgLabelSchemaBuildDate       time.Time `json:"org.label-schema.build-date"`
	OrgOpencontainersImageCreated time.Time `json:"org.opencontainers.image.created"`
}

func NewWorkloads(images []Image, services Services) []Workload {
	workloads := []Workload{}
	for _, v := range services.Services {
		var w Workload
		w.ID = v.ID
		w.ReadOnly = v.ReadOnly
		w.Status = v.Status
		w.Automated = v.Automated
		for _, i := range images {
			if i.ID == v.ID {
				w.Containers = i.Containers
				break
			}
		}
		w.Policies.Automated = v.Policies.Automated
		w.Policies.TagChartImage = v.Policies.TagChartImage

		for i, containers := range w.Containers {
			w.Containers[i].Available = filterAvailableImages(containers.Available, v.Policies.TagChartImage)
		}

		workloads = append(workloads, w)
	}

	return workloads
}

func filterAvailableImages(availableContainers []Available, policy string) []Available {
	availableImages := []Available{}
	for _, av := range availableContainers {
		filter := strings.TrimSuffix(strings.TrimPrefix(policy, "glob:"), "*")
		res := strings.Split(av.ID, ":")
		if len(res) != 2 {
			log.Printf("error can't create available images for %s ", av.ID)
			return nil
		}
		match, _ := regexp.MatchString(filter, res[1])
		if match {
			availableImages = append(availableImages, av)
		}
	}
	return availableImages
}
