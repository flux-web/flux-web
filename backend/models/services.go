package models

import (
	"encoding/json"

	"github.com/astaxie/beego/httplib"
)

type Services struct {
	Services []Service
}
type Service struct {
	ID         string `json:"ID"`
	Containers []Container `json:"Containers"`
	ReadOnly  string `json:"ReadOnly"`
	Status    string `json:"Status"`
	Automated bool   `json:"Automated"`
	Policies  struct {
		Automated     string `json:"automated"`
		TagChartImage string `json:"tag.chart-image"`
	} `json:"Policies"`
}

func NewServices(url string) (Services, error) {
	var s Services
	var servicesArray []Service

	res, err := httplib.Get(url).Debug(true).Bytes()
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(res, &servicesArray)
	s.Services = servicesArray

	return s, err
}

func (s *Services) GetStatusAutomateByWorkload(workload string) bool {
	for _, v := range s.Services {
		if v.ID == workload {
			return v.Automated
		}
	}
	return false
}

func (s *Services) WantedImageAlreadyDeployed(workload, image string) bool {
	for _, v := range s.Services {
		if v.ID == workload {
			for _, i := range v.Containers {
				if i.Current.ID == image {
					return true
				}
			}
		}
	}
	return false
}
