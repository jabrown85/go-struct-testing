package buildpack

import (
	"encoding/json"
)

type Layer interface {
	LayerName() string
	V2Behavior() bool
}

type layerv1 struct {
	Name string `json:"name"`
}

func (l layerv1) LayerName() string {
	return l.Name
}

func (l layerv1) V2Behavior() bool {
	return false
}

type layerv2 struct {
	Name   string `json:"name"`
	V2Flag *bool  `json:"v2flag"`
}

func (l layerv2) LayerName() string {
	return l.Name
}

func (l layerv2) V2Behavior() bool {
	if l.V2Flag == nil {
		return true
	}
	return *l.V2Flag
}

func Parse(jsonString string, bpVersion int) Layer {
	if bpVersion == 1 {
		var layerTOML layerv1
		json.Unmarshal([]byte(jsonString), &layerTOML)
		// I think it would be easier to force schemas using https://golang.org/pkg/encoding/json/#Decoder.DisallowUnknownFields
		// if we wanted to be that strict
		return layerTOML
	}
	if bpVersion == 2 {
		var layerTOML layerv2
		json.Unmarshal([]byte(jsonString), &layerTOML)
		return layerTOML
	}
	return nil
}
