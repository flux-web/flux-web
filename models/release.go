
package models

import (
	"encoding/json"
)

//type Release struct {
//	Cause struct {
//		Message string `json:"Message"`
//		User    string `json:"User"`
//	} `json:"Cause"`
//	Spec struct {
//		ContainerSpecs struct {
//			DefaultDeployment []struct {
//				Container string `json:"Container"`
//				Current   string `json:"Current"`
//				Target    string `json:"Target"`
//			} `json:"default:deployment"`
//		} `json:"ContainerSpecs"`
//		Kind           string `json:"Kind"`
//		SkipMismatches bool   `json:"SkipMismatches"`
//	} `json:"Spec"`
//	Type string `json:"Type"`
//}
	
type Release struct {
	Cause struct {
		Message string `json:"Message"`
		User    string `json:"User"`
	} `json:"Cause"`
	Spec struct {
		ContainerSpecs struct {
			DefaultDeploymentFluxWeb []struct {
				Container string `json:"Container"`
				Current   string `json:"Current"`
				Target    string `json:"Target"`
			} `json:"default:deployment/flux-web"`
		} `json:"ContainerSpecs"`
		Kind           string `json:"Kind"`
		SkipMismatches bool   `json:"SkipMismatches"`
	} `json:"Spec"`
	Type string `json:"Type"`
}

func NewRelese(data []byte) (Release, error) {
	var r Release
	err := json.Unmarshal(data, &r)
	return r, err
}