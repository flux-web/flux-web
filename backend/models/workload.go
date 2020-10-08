package models

import (
	"time"
)

type Workload struct {
	ID         string      `json:"ID"`
	Containers []Container `json:"Containers"`
	ReadOnly   string      `json:"ReadOnly"`
	Status     string      `json:"Status"`
	Automated  bool        `json:"Automated"`
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

func NewWorkloads(images []Image, services []Service) []Workload {

	var workloadsMap = make(map[string]*Workload)

	// generate a map of Workloads from a list of Services
	for i := range services {
		w := new(Workload)
		w.ID = services[i].ID
		w.ReadOnly = services[i].ReadOnly
		w.Status = services[i].Status
		w.Automated = services[i].Automated
		workloadsMap[services[i].ID] = w
	}

	// enrich the Workloads with data from a list of Images
	for i := range images {
		w := workloadsMap[images[i].ID]

		containerCount := len(images[i].Containers)
		w.Containers = make([]Container, containerCount, containerCount)
		for j := range images[i].Containers {

			w.Containers[j].Name = images[i].Containers[j].Name
			w.Containers[j].Current.ID = images[i].Containers[j].Current.ID
			w.Containers[j].Current.Digest = images[i].Containers[j].Current.Digest
			w.Containers[j].Current.ImageID = images[i].Containers[j].Current.ImageID
			w.Containers[j].Current.Labels.OrgLabelSchemaBuildDate = images[i].Containers[j].Current.Labels.OrgLabelSchemaBuildDate
			w.Containers[j].Current.Labels.OrgOpencontainersImageCreated = images[i].Containers[j].Current.Labels.OrgOpencontainersImageCreated
			w.Containers[j].Current.CreatedAt = images[i].Containers[j].Current.CreatedAt
			w.Containers[j].Current.LastFetched = images[i].Containers[j].Current.LastFetched
			w.Containers[j].LatestFiltered.ID = images[i].Containers[j].LatestFiltered.ID
			w.Containers[j].LatestFiltered.Digest = images[i].Containers[j].LatestFiltered.Digest
			w.Containers[j].LatestFiltered.ImageID = images[i].Containers[j].LatestFiltered.ImageID
			w.Containers[j].LatestFiltered.Labels.OrgLabelSchemaBuildDate = images[i].Containers[j].LatestFiltered.Labels.OrgLabelSchemaBuildDate
			w.Containers[j].LatestFiltered.Labels.OrgOpencontainersImageCreated = images[i].Containers[j].LatestFiltered.Labels.OrgOpencontainersImageCreated
			w.Containers[j].LatestFiltered.CreatedAt = images[i].Containers[j].LatestFiltered.CreatedAt
			w.Containers[j].LatestFiltered.LastFetched = images[i].Containers[j].LatestFiltered.LastFetched

			availableCount := len(images[i].Containers[j].Available)
			w.Containers[j].Available = make([]Available, availableCount, availableCount)
			for k := range images[i].Containers[j].Available {
				w.Containers[j].Available[k].ID = images[i].Containers[j].Available[k].ID
				w.Containers[j].Available[k].Digest = images[i].Containers[j].Available[k].Digest
				w.Containers[j].Available[k].ImageID = images[i].Containers[j].Available[k].ImageID
				w.Containers[j].Available[k].Labels.OrgLabelSchemaBuildDate = images[i].Containers[j].Available[k].Labels.OrgLabelSchemaBuildDate
				w.Containers[j].Available[k].Labels.OrgOpencontainersImageCreated = images[i].Containers[j].Available[k].Labels.OrgOpencontainersImageCreated
				w.Containers[j].Available[k].CreatedAt = images[i].Containers[j].Available[k].CreatedAt
				w.Containers[j].Available[k].LastFetched = images[i].Containers[j].Available[k].LastFetched
			}

			w.Containers[j].AvailableImagesCount = images[i].Containers[j].AvailableImagesCount
			w.Containers[j].FilteredImagesCount = images[i].Containers[j].FilteredImagesCount
		}
	}

	// convert map to array
	workloadsList := []Workload{}
	for _, value := range workloadsMap {
		workloadsList = append(workloadsList, *value)
	}

	return workloadsList
}
