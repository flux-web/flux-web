package models

import (
	"encoding/json"
	"fmt"
)

type DeploySettings struct {
	Add    Automated `json:"add,omitempty"`
	Remove Automated `json:"remove,omitempty"`
}

type Automated struct {
	Automated string `json:"automated,omitempty"`
}

type AutomatedRequest struct {
	Type  string `json:"Type"`
	Cause struct {
		Message string `json:"Message"`
		User    string `json:"User"`
	} `json:"Cause"`
	Spec map[string]DeploySettings `json:"Spec"`
}

type DynamicAutomatedRequest struct {
	Workload string `json:"-"`
	Payload  AutomatedRequest
}

func (i DynamicAutomatedRequest) MarshalJSON() ([]byte, error) {
	orig, err := json.Marshal(i.Payload)
	if err != nil {
		return nil, err
	}

	var tmp interface{}
	err = json.Unmarshal(orig, &tmp)

	b := tmp.(map[string]interface{})

	b[spec].(map[string]interface{})[i.Workload] = b[spec].(map[string]interface{})[placeHolder]
	delete(b[spec].(map[string]interface{}), placeHolder)

	return json.Marshal(b)
}

func (this *ReleaseRequest) GetAutomatedRequestJSON(fluxUser string) ([]byte, error) {
	req := DynamicAutomatedRequest{}
	req.Workload = this.Workload

	d := DeploySettings{}
	if this.Automated {
		d.Add.Automated = fmt.Sprintf("%t", this.Automated)
	} else {
		d.Remove.Automated = fmt.Sprintf("%t", this.Automated)
	}
	spec := make(map[string]DeploySettings)
	spec[placeHolder] = d
	req.Payload.Spec = spec
	req.Payload.Type = "policy"
	req.Payload.Cause.User = fluxUser

	return json.Marshal(req)
}
