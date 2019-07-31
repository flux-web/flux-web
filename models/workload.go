package models

import (
//"encoding/json"
	"time"
)

type Workload struct {
	ID         string `json:"ID"`
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

//func NewWorkloads(images []Image, services []Service) ([]Workload, error) {
//	var w []Workload
//	err := json.Unmarshal(data, &w)
//	return w, er
//	return []Workload
//}

