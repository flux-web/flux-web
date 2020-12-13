package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/nsf/jsondiff"
	"github.com/stretchr/testify/assert"
)

func TestDynamicRequestPayload_MarshalJSON(t *testing.T) {
	static, err := ioutil.ReadFile("testdata/dynamicRequestPayload.json")
	if err != nil {
		t.Fatal(err)
	}

	req := DynamicRequestPayload{}
	req.Workload = "lorem:helmrelease/foobar"

	c := make(map[string][]ContainerUpdate)
	c[placeHolder] = []ContainerUpdate{{
		Container: "chart-image",
		Current:   "my-registry.io/images/foobar:7c5106v",
		Target:    "my-registry.io/images/foobar:42g3z6c",
	}}

	req.Payload.Cause.User = "Flux"
	req.Payload.Type = TypeContainers
	req.Payload.Spec.Kind = kindExecute
	req.Payload.Spec.SkipMismatches = false
	req.Payload.Spec.ContainerSpecs = c
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	diffOpts := jsondiff.DefaultConsoleOptions()
	res, diff := jsondiff.Compare(data, static, &diffOpts)
	assert.Equal(t, jsondiff.FullMatch, res, "compare marshaled json with static json")
	if res != jsondiff.FullMatch {
		fmt.Println(diff)
	}
}
