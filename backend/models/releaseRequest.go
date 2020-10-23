package models

import (
	"encoding/json"
)

type ReleaseRequest struct {
	Workload  string `json:"Workload"`
	Container string `json:"Container"`
	Current   string `json:"Current"`
	Target    string `json:"Target"`
	Automated bool   `json:"Automated"`
}

type ContainerUpdate struct {
	Container string `json:"Container"`
	Current   string `json:"Current"`
	Target    string `json:"Target"`
}

type RequestPayload struct {
	Cause struct {
		Message string `json:"Message"`
		User    string `json:"User"`
	} `json:"Cause"`
	Spec struct {
		ContainerSpecs map[string][]ContainerUpdate `json:"ContainerSpecs"`
		Kind           string                       `json:"Kind"`
		SkipMismatches bool                         `json:"SkipMismatches"`
	} `json:"Spec"`
	Type string `json:"Type"`
}

type DynamicRequestPayload struct {
	Workload string `json:"-"`
	Payload  RequestPayload
}

func (i DynamicRequestPayload) MarshalJSON() ([]byte, error) {
	orig, err := json.Marshal(i.Payload)
	if err != nil {
		return nil, err
	}

	var tmp interface{}
	err = json.Unmarshal(orig, &tmp)

	b := tmp.(map[string]interface{})

	b[spec].(map[string]interface{})[containerSpecs].(map[string]interface{})[i.Workload] = b[spec].(map[string]interface{})[containerSpecs].(map[string]interface{})[placeHolder]
	delete(b[spec].(map[string]interface{})[containerSpecs].(map[string]interface{}), placeHolder)

	return json.Marshal(b)
}

func NewReleseRequest(data []byte) (ReleaseRequest, error) {
	var r ReleaseRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (this *ReleaseRequest) GetReleaseRequestJSON(fluxUser string) ([]byte, error) {
	req := DynamicRequestPayload{}
	req.Workload = this.Workload

	c := make(map[string][]ContainerUpdate)
	c[placeHolder] = []ContainerUpdate{{
		Container: this.Container,
		Current:   this.Current,
		Target:    this.Target,
	}}

	req.Payload.Cause.User = fluxUser
	req.Payload.Type = TypeContainers
	req.Payload.Spec.Kind = kindExecute
	req.Payload.Spec.SkipMismatches = false
	req.Payload.Spec.ContainerSpecs = c

	return json.Marshal(req)
}
