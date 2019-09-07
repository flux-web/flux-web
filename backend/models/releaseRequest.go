package models

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

type ReleaseRequest struct {
	Workload  string `json:"Workload"`
	Container string `json:"Container"`
	Current   string `json:"Current"`
	Target    string `json:"Target"`
}

func NewReleseRequest(data []byte) (ReleaseRequest, error) {
	var r ReleaseRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (this *ReleaseRequest) GetReleaseRequestJSON() []byte {
	spec := "\"" + this.Workload + "\":[{\"Container\":\"" + this.Container + "\",\"Current\":\"" + this.Current + "\",\"Target\":\"" + this.Target + "\"}]"
	releaseRequest := "{\"Cause\":{\"Message\":\"\", \"User\":\"Flux-web\"},\"Spec\":{\"ContainerSpecs\":{" + spec + "},\"Kind\":\"execute\",\"SkipMismatches\":true},\"Type\":\"containers\"}"
	var l = logs.GetLogger()
	l.Println("****************************")
	l.Println(releaseRequest)
	l.Println("****************************")
	return []byte(releaseRequest)
}
