package models

type ReleaseRequestFlux struct {
	Cause struct {
		Message string `json:"Message"`
		User    string `json:"User"`
	} `json:"Cause"`
	Spec struct {
		ContainerSpecs map[string]Helmrelease `json:"ContainerSpecs"`
		Kind           string                 `json:"Kind"`
		SkipMismatches bool                   `json:"SkipMismatches"`
	} `json:"Spec"`
	Type string `json:"Type"`
}

type Helmrelease struct {
	Container string `json:"Container"`
	Current   string `json:"Current"`
	Target    string `json:"Target"`
}

//{
//	"Cause": {
//		"Message": "",
//		"User": "Flux-web"
//	},
//	"Spec": {
//		"ContainerSpecs": {
//			"red-stg:helmrelease/red-cmp-service": [
//				{
//				"Container": "chart-image",
//				"Current": "spring-docker.jfrog.io/red/red-cmp-service:e65f584",
//				"Target": "spring-docker.jfrog.io/red/red-cmp-service:412c991"
//				}
//			]
//		},
//		"Kind": "execute",
//		"SkipMismatches": false
//	},
//	"Type": "containers"
//}

//func ReleaseRequestFlux(data []byte) (ReleaseRequest, error) {
//	var r ReleaseRequest
//	err := json.Unmarshal(data, &r)
//	return r, err
//}
