package models

import (
	"encoding/json"
)

type Image struct {
	ID         string `json:"ID"`
	Containers []Container `json:"Containers"`
}

func NewImages(data []byte) ([]Image, error) {
	var i []Image
	err := json.Unmarshal(data, &i)
	return i, err
}
